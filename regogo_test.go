package regogo

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		query  string
		expect Result
	}{
		{
			query:  "input.hello",
			expect: Result{Value: "world"},
		},
		{
			query:  "input.items[0].val",
			expect: Result{Value: json.Number("1")},
		},
		{
			query:  "input.foo.bool",
			expect: Result{Value: true},
		},
		{
			query:  "input.items[_].val",
			expect: Result{Value: []interface{}{json.Number("1"), json.Number("2"), []interface{}{json.Number("1"), json.Number("2"), json.Number("3")}}},
		},
		{
			query:  "count(input.items)",
			expect: Result{Value: json.Number("3")},
		},
		{
			query:  "{\"foo\": [input.items[0].val, input.items[1].val]}",
			expect: Result{Value: map[string]interface{}{"foo": []interface{}{json.Number("1"), json.Number("2")}}},
		},
		{
			query:  "[ v | walk(input, [_, v]); is_number(v) ]",
			expect: Result{Value: []interface{}{json.Number("1"), json.Number("2"), json.Number("1"), json.Number("2"), json.Number("3")}},
		},
		{
			query:  "{ v | walk(input, [_, v]); is_number(v) }",
			expect: Result{Value: []interface{}{json.Number("1"), json.Number("2"), json.Number("3")}},
		},
		{
			query:  "i:=0; x:=input.items[i].val; input.items[x].val",
			expect: Result{Value: json.Number("2")},
		},
	}

	testFilePath := "./test/helloworld.json"
	inputBytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("can't open test JSON file: %s", testFilePath)
	}
	input := string(inputBytes)

	for _, test := range tests {
		res, err := Get(input, test.query)
		if err != nil {
			t.Errorf("Get failed: %v", err)
		}
		if !reflect.DeepEqual(res, test.expect) {
			t.Errorf("test failed: %s\nwant %#v, have %#v", test.query, test.expect, res)
		}
	}
}
