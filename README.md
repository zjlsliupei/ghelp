# ghelp
封装go语言常用助手函数

### 安装
go get github.com/zjlsliupei/ghelp

### 编码类方法列表
```go
// json
data := map[string]interface{}{
    "name": "ghelp"
}
encodeStr := ghelp.JsonEncode(data)
decodeMap := ghelp.JsonDecode(encodeStr)

// md5 sha1
str := "1111"
encodeMd5 = ghelp.Md5(str) 
encodeSha1 = ghelp.Sha1(str) 
fmt.Println(encodeMd5, encodeSha1)
// md5:b59c67bf196a4758191e42f76670ceba 
// sha1:011c945f30ce2cbafc452f39840f025693339c42

// base64
data := "hello world"
encodeStr := ghelp.Base64Encode(data)
decodeStr := ghelp.Base64Decode(encodeStr)
fmt.Println(encodeStr, decodeStr)
// encode: VsbG8gd29ybGQ=
// decode: hello world
```


