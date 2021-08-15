package tool

import jsoniter "github.com/json-iterator/go"

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	bytes, err := jsoniter.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}
