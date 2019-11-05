// json常用编码、解码操作

package encode

import (
	"encoding/json"
)

// json编码操作
func JsonEncode(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// json解码操作
func JsonDecode(jsonStr string) interface{} {
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		return nil
	}
	return maps
}
