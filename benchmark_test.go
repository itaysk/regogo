package regogo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Jeffail/gabs/v2"
	"github.com/tidwall/gjson"
)

var input string

func TestMain(m *testing.M) {

	testFilePath := "./test/helloworld.json"
	inputBytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		fmt.Printf("can't open test JSON file: %s\n", testFilePath)
		return
	}
	input = string(inputBytes)

	os.Exit(m.Run())
}

func BenchmarkRegogo(b *testing.B) {
	tests := []string{
		"input.hello",
		"input.foo.bool",
		"input.items[0].val",
	}

	for i := 0; i < b.N; i++ {
		for _, t := range tests {
			_, err := Get(input, t)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkGjson(b *testing.B) {
	tests := []string{
		"hello",
		"foo.bool",
		"items.0.val",
	}

	for i := 0; i < b.N; i++ {
		for _, t := range tests {
			res := gjson.Get(input, t)
			if !res.Exists() {
				b.Error()
			}
		}
	}
}

func BenchmarkGabs(b *testing.B) {
	tests := []string{
		"hello",
		"foo.bool",
		"items.0.val",
	}
	inputBytes := []byte(input)

	for i := 0; i < b.N; i++ {
		for _, t := range tests {
			c, err := gabs.ParseJSON(inputBytes)
			if err != nil {
				b.Error(err)
			}
			res := c.Path(t)
			if res.Data() == nil {
				b.Error()
			}
		}

	}
}
