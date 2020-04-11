package regogo

import (
	"context"
	"encoding/json"

	"github.com/open-policy-agent/opa/rego"
)

func Get(in string, query string) (Result, error) {
	ctx := context.TODO()
	regoQuery, err := rego.New(
		rego.Query(query),
	).PrepareForEval(ctx)
	if err != nil {
		return Result{}, err
	}
	var inMap map[string]interface{}
	err = json.Unmarshal([]byte(in), &inMap)
	if err != nil {
		return Result{}, err
	}
	resset, err := regoQuery.Eval(ctx, rego.EvalInput(inMap))
	if err != nil {
		return Result{}, err
	}

	rescount := len(resset)
	tmpres := make([]interface{}, rescount)
	for i, res := range resset {
		exp := res.Expressions
		tmpres[i] = exp[len(exp)-1].Value
	}
	if rescount == 1 {
		return Result{Value: tmpres[0]}, nil
	}
	return Result{Value: tmpres}, nil
}
