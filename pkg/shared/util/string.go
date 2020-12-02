package util

import "encoding/json"

func Stringify(data interface{}) string {
	dataByte, _ := json.MarshalIndent(data, "", " ")
	return string(dataByte)
}
