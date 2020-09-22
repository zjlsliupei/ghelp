package ghelp

import (
	"testing"
)

type TestData struct {
	Name string
	Sex  string
	Age  byte
	Love []string
}

var (
	jsonStr = `{"age":18,"love":["basketball","football"],"name":"liupei","sex":"man"}`
)

func TestJsonEncode(t *testing.T) {
	// map
	maps := map[string]interface{}{
		"name": "liupei",
		"sex":  "man",
		"age":  18,
		"love": []string{
			"basketball", "football",
		},
	}
	jsonStr := JsonEncode(maps)
	if jsonStr == "" {
		t.Errorf("map to json error!")
	}
	// struct
	newPerson := TestData{
		Name: "liupei",
		Sex:  "man",
		Age:  18,
		Love: []string{
			"basketball", "football",
		},
	}
	jsonStr2 := JsonEncode(newPerson)
	if jsonStr2 == "" {
		t.Errorf("struct to json error!")
	}
}

func TestJsonDecode(t *testing.T) {
	maps := JsonDecode(jsonStr)
	if maps == nil {
		t.Error("json_decode error")
	}
}

func TestMd5(t *testing.T) {
	md5Str := Md5("1111")
	if md5Str != "b59c67bf196a4758191e42f76670ceba" {
		t.Error("md5 error")
	}
}

func TestSha1(t *testing.T) {
	sha1Str := Sha1("1111")
	if sha1Str != "011c945f30ce2cbafc452f39840f025693339c42" {
		t.Error("sha1 error")
	}
}

func TestBase64Encode(t *testing.T) {
	str := Base64Encode("hello world")
	if str != "aGVsbG8gd29ybGQ=" {
		t.Error("base64Encode error")
	}
}

func TestBase64Decode(t *testing.T) {
	str := Base64Decode("aGVsbG8gd29ybGQ=")
	if str != "hello world" {
		t.Error("base64Decode error")
	}
}
