package controller

import (
	"fmt"
	"regexp"
	"testing"
)

func parseRequestRegexp(raw string, typeIdRegexp *regexp.Regexp) *Request {
	found := typeIdRegexp.FindAllString(raw, 2)
	if len(found) < 2 {
		return nil
	}
	return &Request{
		Type: RequestType(found[0][10:]),
		Id:   found[1][8:],
	}
}

const requestRaw = "{\"type\":\"aAbcdefgzZ\",\"id\":\"cb077db43b627bb7\",\"key1\":\"qwertyuiop[asdfghjkl;\",\"key2\":true,\"key3\":1234567890,\"obj\":{\"innerKey\":\"1234abcd\"}}"

func TestParseRequestManual(t *testing.T) {
	if parseRequestManual(requestRaw) == nil {
		t.FailNow()
	}
	if parseRequestManual("{\"type\":\"\",\"id\":\"\",\"key1\":\"qwertyuiop[asdfghjkl;\"}") == nil {
		t.FailNow()
	}
	if parseRequestManual("{\"type\":12345,\"id\":\"cb077db43b627bb7\",\"key1\":\"qwertyuiop[asdfghjkl;\"}") != nil {
		t.FailNow()
	}
	if parseRequestManual("{\"type\":\"join\",\"id\":\"0123456789abcdef0\",\"key1\":\"qwertyuiop[asdfghjkl;\"}") != nil {
		t.FailNow()
	}
	if parseRequestManual("{\"type\":") != nil {
		t.FailNow()
	}
}

func BenchmarkParseRequestManual(b *testing.B) {
	fmt.Println(parseRequestManual(requestRaw))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseRequestManual(requestRaw)
	}
}

func BenchmarkParseRequestRegexp(b *testing.B) {
	typeIdRegexp := regexp.MustCompile(`({"type":"[a-zA-Z0-9]+)|((,"id":")[a-zA-Z0-9]+)`)
	fmt.Println(parseRequestRegexp(requestRaw, typeIdRegexp))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseRequestRegexp(requestRaw, typeIdRegexp)
	}
}

func BenchmarkParseRequestJson(b *testing.B) {
	fmt.Println(parseRequest(&Request{}, requestRaw))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseRequest(&Request{}, requestRaw)
	}
}
