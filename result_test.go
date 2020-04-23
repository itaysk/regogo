package regogo

import (
	"encoding/json"
	"testing"
)

func TestMarshalText(t *testing.T) {
	tests := []struct {
		input  Result
		expect string
	}{
		{
			input:  Result{Value: "world"},
			expect: "world",
		},
		{
			input:  Result{Value: true},
			expect: "true",
		},
		{
			input:  Result{Value: json.Number("1")},
			expect: "1",
		},
		{
			input:  Result{Value: []interface{}{json.Number("1"), "foo", true, []interface{}{json.Number("2"), "bar", false}}},
			expect: "[1, foo, true, [2, bar, false]]",
		},
		{
			input: Result{Value: map[string]interface{}{"foo": json.Number("1"), "bar": map[string]interface{}{"baz": true}}},
			// maps are sorted
			expect: "{bar: {baz: true}, foo: 1}",
		},
	}
	for _, test := range tests {
		res, err := test.input.MarshalText()
		if err != nil {
			t.Errorf("MarshalText failed: %v", err)
		}
		if string(res) != test.expect {
			t.Errorf("test failed: %+v\nwant %s, have %s", test.input, test.expect, res)
		}
	}
}
