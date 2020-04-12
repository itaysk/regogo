This is a Go package for querying arbitrary JSON documents using a query string. Actually, it's just a wrapper over Rego (the [Open Policy Agent](https://www.openpolicyagent.org/) engine) which does all the work. Nevertheless, this package tries to make Rego more accessible for general purpose JSON querying.

# Quickstart

```go
import "github.com/itaysk/regogo"
...
myjson := "{...}"
res, err = regogo.Get(myjson, "myquery")
if err != nil {
	...
}
fmt.Println(res.String())
```

# Examples

Consider the input JSON document:

```json
{
	"hello": "world",
	"items" : [
		{"val": 1},
		{"val": 2},
		{"val": [1,2,3]}
	],
	"foo": {
		"bar": "baz",
		"bool": true
	}
}
```

The following queries will produce the following results:

query|result
---|---
`input.hello` | `"world"`
`input.items[0].val` | `1`
`count(input.items)` | `3`
`[ v \| walk(input, [_, v]); is_number(v) ]` |  `[1, 2, 1, 2, 3]`

You can find more examples in the tests for the `Get` function: [regogo_test.go](regogo_test.go)

All of this is standard Rego, but for most basic queries you can even guess your way into a working query without learning Rego.

# API Documentation

## Query

The query is any valid Rego query, the JSON document you are querying will be represented in the query by the variable `input`.

Rego can return complex output containing multiple results, and regogo aspires to simplify that. Therefore the return value for `Get` is built by picking the value of the last expression in each result in the result set. If there are multiple values to return, the return value is an array. This convention allows for multi part queries, while maintaining a simple API.  
**You don't need to worry about any of this if your queries are simple one liners, like `input.array[i].sub`.**

For a tutorial on Rego, see [here](https://www.openpolicyagent.org/docs/latest/#rego).  
For a complete reference of Rego, see [here](https://www.openpolicyagent.org/docs/latest/policy-language).

## Result

The returned object is a `regogo.Result` type that just encapsulates the returned value and provides convenience methods to use it, such as getting the value as string, number, bool, array, etc.

# Background

I think working with JSON in Go is a pain, and none of the available libraries did it right (closest was gjson).  
After using OPA for another project, I noticed the similarity between rego and other JSON libraries. The messaging was completely different but the technology was similar.  
I tries to use rego instead of gjson and it was OK, but there was too much Rego boilerplate to get a simple value from a document, so I created this wrapper to simplify that. 
By making some assumptions and conventions we can create a simple interface to query JSON documents using Rego.

## Is this for real?

Good question, I'm unsure myself. On one hand it's too simple to feel like a "library", but on the other, it is useful. I guess we'll know based on your feedback (or lack of). In any case, please don't be wrongfully impressed by the fact that I put up this fancy readme, consider it more like an accompanying blog post.