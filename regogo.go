package regogo

import (
	"context"
	"encoding/json"

	"github.com/open-policy-agent/opa/rego"
)

type Regogo struct {
	query     string
	evaluator rego.PreparedEvalQuery
}

func (rg Regogo) Get(input string) (Result, error) {
	var inMap map[string]interface{}
	err := json.Unmarshal([]byte(input), &inMap)
	if err != nil {
		return Result{}, err
	}
	resset, err := rg.evaluator.Eval(context.TODO(), rego.EvalInput(inMap))
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

func New(query string) (*Regogo, error) {
	regoQuery, err := rego.New(
		rego.Query(query),
	).PrepareForEval(context.TODO())
	if err != nil {
		return nil, err
	}
	return &Regogo{query: query, evaluator: regoQuery}, nil
}

func Get(input string, query string) (Result, error) {
	rg, err := New(query)
	if err != nil {
		return Result{}, err
	}
	return rg.Get(input)
}
