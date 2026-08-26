package util

import (
	"github.com/lithammer/shortuuid/v5"
)

func ShortUUID() string {
	return shortuuid.New()
}

func ShortUUIDUnique(isUnique func(value string) bool) string {
	value := shortuuid.New()
	for !isUnique(value) {
		value = shortuuid.New()
	}
	return value
}
