package util

import (
	"encoding/json"
	"fmt"
)

func JsonToObject(object interface{}, config string) interface{} {
	err := json.Unmarshal([]byte(config), object)
	if err != nil {
		fmt.Printf("Can't parse %T: %v", object, err)
	}
	return object
}
