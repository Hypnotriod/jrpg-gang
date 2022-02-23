package util

import (
	"encoding/json"
)

func JsonToObject(object interface{}, config string) (interface{}, error) {
	return object, json.Unmarshal([]byte(config), object)
}

func ObjectToJson(object interface{}) string {
	if marshalled, err := json.Marshal(object); err == nil {
		return string(marshalled)
	}
	return ""
}
