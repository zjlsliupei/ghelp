package ghelp

import (
	"fmt"
	"reflect"
)

// InArray 判断元素是否在数组中
// 支持string,int,int8,int16,int32,int64,float32,float64,bool类型
func InArray(item interface{}, items interface{}) bool {
	switch itemsValue := items.(type) {
	case []string:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int8:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int16:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int32:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int64:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []float32:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []float64:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []bool:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	}
	fmt.Println(reflect.TypeOf(item), reflect.TypeOf(items))
	return false
}
