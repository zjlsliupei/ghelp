package excel_column

import (
	"strings"
)

// FirstColName
// 获取首列名称
func FirstColName() string {
	return "A"
}

// LastColName
// 获取尾列名称
func LastColName() string {
	return "XFD"
}

// NextColName
// 获取下一列名称
func NextColName(colName string) string {
	arr := splitWord(colName)
	if LastColName() == colName {
		return ""
	}
	nextColName := after(arr[0], arr[1], arr[2])
	return nextColName
}

// NextColName
// 获取上一列名称
func PreColName(colName string) string {
	arr := splitWord(colName)
	if FirstColName() == colName {
		return ""
	}
	nextColName := before(arr[0], arr[1], arr[2])
	return nextColName
}

func after(first, second, third string) string {
	if !isEnd(third) {
		return first + second + convertToString(convertToInt(third)+1)
	}
	// 最后一位是Z, 前一位不是Z
	if !isEnd(second) {
		if second == "" {
			return first + "AA"
		} else {
			return first + convertToString(convertToInt(second)+1) + "A"
		}
	}
	// 中间位是Z, 前一位不是Z
	if !isEnd(first) {
		if first == "" {
			return "A" + "A" + "A"
		} else {
			return convertToString(convertToInt(first)+1) + "A" + "A"
		}

	}
	return ""
}

func before(first, second, third string) string {
	if !isFirst(third) {
		return first + second + convertToString(convertToInt(third)-1)
	}
	// 最后一位是A, 前一位不是A
	if !isFirst(second) {
		if second == "" {
			return ""
		} else {
			return first + convertToString(convertToInt(second)-1) + "Z"
		}
	}
	// 中间位是A, 前一位不是A
	if !isFirst(first) {
		if first == "" {
			return "Z"
		} else {
			return convertToString(convertToInt(first)-1) + "Z" + "Z"
		}

	}
	return ""
}

func isEnd(word string) bool {
	if word == "z" || word == "Z" {
		return true
	} else {
		return false
	}
}

func isFirst(word string) bool {
	if word == "a" || word == "A" {
		return true
	} else {
		return false
	}
}

func convertToInt(word string) int {
	return int([]byte(word)[0])
}

func convertToString(wordIndex int) string {
	return string(wordIndex)
}

func splitWord(v string) []string {
	arr := strings.Split(v, "")
	count := len(arr)
	if count == 0 {
		arr = []string{"", "", ""}
	} else if count == 1 {
		arr = append([]string{"", ""}, arr...)
	} else if count == 2 {
		arr = append([]string{""}, arr...)
	}
	return arr
}
