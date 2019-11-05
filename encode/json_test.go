package encode

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
