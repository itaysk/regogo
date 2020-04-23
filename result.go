package regogo

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func (r Result) MarshalText() ([]byte, error) {
	switch v := r.Value.(type) {
	case string:
		return []byte(v), nil
	case bool:
		return []byte(strconv.FormatBool(v)), nil
	case fmt.Stringer: //will also catch json.Number
		return []byte(v.String()), nil
	case []interface{}:
		var b strings.Builder
		b.WriteRune('[')
		numsep := len(v) - 1
		for i, vv := range v {
			vvText, err := Result{Value: vv}.MarshalText()
			if err != nil {
				continue
			}
			_, err = b.Write(vvText)
			if err != nil {
				continue
			}
			if i < numsep {
				b.WriteString(", ")
			}
		}
		b.WriteRune(']')
		return []byte(b.String()), nil
	case map[string]interface{}:
		var b strings.Builder
		b.WriteRune('{')
		numsep := len(v) - 1
		//sort map to ensure consistent results
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			vv := v[k]
			vvText, err := Result{Value: vv}.MarshalText()
			if err != nil {
				continue
			}
			_, err = b.WriteString(fmt.Sprintf("%s: %s", k, vvText))
			if err != nil {
				continue
			}
			if i < numsep {
				b.WriteString(", ")
			}
		}
		b.WriteRune('}')
		return []byte(b.String()), nil
	default:
		return []byte{}, nil
	}
}
