package util

import (
	"encoding/json"
)

func JsonToObject(object interface{}, config string) (interface{}, bool) {
	err := json.Unmarshal([]byte(config), object)
	return object, err == nil
}

func ObjectToJson(object interface{}) string {
	if marshalled, err := json.Marshal(object); err == nil {
		return string(marshalled)
	}
	return ""
}
