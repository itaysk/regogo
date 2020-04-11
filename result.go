package regogo

import (
	"encoding/json"
)

type Result struct {
	Value interface{}
}

func (r Result) Number() json.Number {
	v, ok := r.Value.(json.Number)
	if !ok {
		return json.Number("0")
	}
	return v
}

func (r Result) Float64() float64 {
	vnum, ok := r.Value.(json.Number)
	if !ok {
		return 0.
	}
	v, err := vnum.Float64()
	if err != nil {
		return 0.
	}
	return v
}

func (r Result) Int64() int64 {
	vnum, ok := r.Value.(json.Number)
	if !ok {
		return 0
	}
	v, err := vnum.Int64()
	if err != nil {
		return 0
	}
	return v
}

func (r Result) Bool() bool {
	v, ok := r.Value.(bool)
	if !ok {
		return false
	}
	return v
}

func (r Result) String() string {
	v, ok := r.Value.(string)
	if !ok {
		return ""
	}
	return v
}

func (r Result) JSON() string {
	vBytes, err := json.Marshal(r.Value)
	if err != nil {
		return ""
	}
	return string(vBytes)
}
