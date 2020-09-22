// json常用编码、解码操作

package ghelp

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// JsonEncode json编码操作
func JsonEncode(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// JsonDecode json解码操作
// 成功返回map, 失败返回nil
func JsonDecode(jsonStr string) map[string]interface{} {
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		return nil
	}
	return maps
}

// Base64Encode base64编码操作
func Base64Encode(data string) string {
	bData := []byte(data)
	return base64.StdEncoding.EncodeToString(bData)
}

// Base64Encode base64解码操作
func Base64Decode(encodeStr string) string {
	bData, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		return ""
	}
	return string(bData)
}

// Md5
func Md5(data string) string {
	h := md5.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1
func Sha1(data string) string {
	h := sha1.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
