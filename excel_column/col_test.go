package excel_column

import (
	"testing"
)

func TestNextColName(t *testing.T) {
	if NextColName("Z") != "AA" {
		t.Error("Z next is not AA")
	}
	if NextColName("A") != "B" {
		t.Error("A next is not B")
	}

	if NextColName("AB") != "AC" {
		t.Error("AB next is not AC")
	}

	if NextColName("AZ") != "BA" {
		t.Error("AZ next is not BA")
	}
	if NextColName("ZZ") != "AAA" {
		t.Error("ZZ next is not AAA")
	}
	if NextColName("ABC") != "ABD" {
		t.Error("ABC next is not ABD")
	}
	if NextColName("ABZ") != "ACA" {
		t.Error("A next is not B")
	}
}

func TestPreColName(t *testing.T) {
	if PreColName("AA") != "Z" {
		t.Error("AA previous is not Z")
	}
	if PreColName("B") != "A" {
		t.Error("B previous is not A")
	}

	if PreColName("AC") != "AB" {
		t.Error("AC previous is not AB")
	}

	if PreColName("BA") != "AZ" {
		t.Error("BA previous is not AZ")
	}
	if PreColName("AAA") != "ZZZ" {
		t.Error("AAA previous is not ZZZ")
	}
	if PreColName("ABD") != "ABC" {
		t.Error("ABD previous is not ABC")
	}
	if PreColName("ACA") != "ABZ" {
		t.Error("ACA previous is not ABZ")
	}
}

func TestFirstColName(t *testing.T) {
	if FirstColName() != "A" {
		t.Error("FirstColName is not A")
	}
}

func TestLastColName(t *testing.T) {
	if LastColName() != "A" {
		t.Error("LastColName is not A")
	}
}
