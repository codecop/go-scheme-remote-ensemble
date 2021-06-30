package utils

import (
	"testing"
)

func TestRemoveLineBreaks(t *testing.T) {
	expected := "a b c"
	actual := removeLineBreaks("a\nb\nc")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}

func TestRemoveSpaces(t *testing.T) {
	expected := "(a b(c(d e)))"
	actual := removeSpaces("(a b(c (d e) ) )")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	expected := "(if(> 2 3)(print \"hello\")(if(< 4 7)(print \"world\")(print \"hello\")))"
	actual := ReadFile("../lisp.txt")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}
