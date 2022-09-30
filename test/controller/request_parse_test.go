package controller

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/util"
	"regexp"
	"testing"
)

func parseRequestJson(data *controller.Request, requestRaw string) *controller.Request {
	r, err := util.JsonToObject(data, requestRaw)
	if err == nil {
		return r.(*controller.Request)
	}
	return nil
}

func parseRequestManual(raw string) *controller.Request {
	var typeStr string
	var idStr string
	r := []byte(raw)
	if len(r) < 10 || r[0] != '{' || r[1] != '"' || r[2] != 't' || r[3] != 'y' || r[4] != 'p' || r[5] != 'e' || r[6] != '"' || r[7] != ':' || r[8] != '"' {
		return nil
	}
	typeBytes := [16]byte{}
	r = r[9:]
	for i, c := range r {
		if c == '"' {
			r = r[i+1:]
			typeStr = string(typeBytes[:i])
			break
		} else if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') || i == len(typeBytes) || i == len(r)-1 {
			return nil
		} else {
			typeBytes[i] = c
		}
	}
	if len(r) < 8 || r[0] != ',' || r[1] != '"' || r[2] != 'i' || r[3] != 'd' || r[4] != '"' || r[5] != ':' || r[6] != '"' {
		return nil
	}
	r = r[7:]
	idBytes := [16]byte{}
	for i, c := range r {
		if c == '"' {
			idStr = string(idBytes[:i])
			break
		} else if !(c >= 'a' && c <= 'h' || c >= '0' && c <= '9') || i == len(idBytes) || i == len(r)-1 {
			return nil
		} else {
			idBytes[i] = c
		}
	}
	return &controller.Request{
		Type: controller.RequestType(typeStr),
		Id:   idStr,
	}
}

func parseRequestRegexp(raw string, typeIdRegexp *regexp.Regexp) *controller.Request {
	found := typeIdRegexp.FindAllString(raw, 2)
	if len(found) < 2 {
		return nil
	}
	return &controller.Request{
		Type: controller.RequestType(found[0][10:]),
		Id:   found[1][8:],
	}
}

const requestRaw = "{\"type\":\"aAbcdefgzZ\",\"id\":\"cb077db43b627bb7\",\"key1\":\"qwertyuiop[asdfghjkl;\",\"key2\":true,\"key3\":1234567890,\"obj\":{\"innerKey\":\"1234abcd\"}}"

func TestParseRequestManual(t *testing.T) {
	if parseRequestManual(requestRaw) == nil {
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
	fmt.Println(parseRequestJson(&controller.Request{}, requestRaw))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseRequestJson(&controller.Request{}, requestRaw)
	}
}
