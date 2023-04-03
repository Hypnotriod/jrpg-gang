package util

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonToObject(obj any, config string) (any, error) {
	return obj, json.Unmarshal([]byte(config), obj)
}

func ObjectToJson(obj any) string {
	if marshalled, err := json.Marshal(obj); err == nil {
		return string(marshalled)
	}
	return ""
}

func ReadJsonFile[T any](obj *T, path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
