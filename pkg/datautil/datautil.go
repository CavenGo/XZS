package datautil

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// CopyStruct 复制结构体，在main函数中开启了jsoiter兼容php模式，因为数字字符串也可以自己动转int，或int也可以自动转string
func CopyStruct(src, dst any) error {
	bytes, _ := json.Marshal(src)
	return json.Unmarshal(bytes, &dst)
}
