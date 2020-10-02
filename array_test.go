package ghelp

import (
	"testing"
)

func TestInArray(t *testing.T) {
	item1 := "a"
	items1 := []string{"a", "b", "c"}
	if !InArray(item1, items1) {
		t.Error("a in not in [a,b,c]")
	}

	item2 := 1
	items2 := []int{1, 2, 3}
	if !InArray(item2, items2) {
		t.Error("1 in not in [1,2,3]")
	}

	item3 := false
	items3 := []bool{false, true}
	if !InArray(item3, items3) {
		t.Error("false in not in [false, true]")
	}
}
