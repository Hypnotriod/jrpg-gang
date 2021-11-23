package util

import (
	"encoding/json"
	"fmt"
)

func JsonToObject(object interface{}, config string) (interface{}, bool) {
	err := json.Unmarshal([]byte(config), object)
	if err != nil {
		fmt.Printf("Can't parse %T: %v", object, err)
	}
	return object, err == nil
}
