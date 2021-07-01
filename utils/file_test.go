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

func TestAddSpacesToBrackets(t *testing.T) {
	expected := "( a b ( c ( d e ) ) )"
	actual := addSpacesToBrackets("(a b(c(d e) ))")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}

func TestRemoveMultipleSpaces(t *testing.T) {
	expected := "(a b(c (d e )) )"
	actual := removeMultipleSpaces("(a       b(c (d e   ))     )")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	expected := "(if (> 2 3) (print \"hello\") (if (< 4 7) ( print \"world\" ) ( print \"hello\" ) ) )"
	actual := ReadFile("../lisp.txt")

	if actual != expected {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}

func isSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSplitTokens(t *testing.T) {
	expected := []string{"(", "if", "(", ">", "2", "3", ")", ")"}
	actual := SplitTokens("( if ( > 2 3 ) )")
	if !isSliceEqual(actual, expected) {
		t.Errorf("Error, expected %s, actual %s", expected, actual)
	}
}
