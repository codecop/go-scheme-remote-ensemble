package parsing

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLineBreaks(t *testing.T) {
	expected := "a b c"
	actual := removeLineBreaks("a\nb\nc")
	assert.Equal(t, expected, actual, fmt.Sprintf("Error, expected %s, actual %s", expected, actual))
}

func TestAddSpacesToBrackets(t *testing.T) {
	expected := "( a b ( c ( d e ) ) )"
	actual := addSpacesToBrackets("(a b(c(d e) ))")
	assert.Equal(t, expected, actual, fmt.Sprintf("Error, expected %s, actual %s", expected, actual))
}

func TestRemoveMultipleSpaces(t *testing.T) {
	expected := "(a b(c (d e )) )"
	actual := removeMultipleSpaces("(a       b(c (d e   ))     )")
	assert.Equal(t, expected, actual, fmt.Sprintf("Error, expected %s, actual %s", expected, actual))
}

func TestTokenize(t *testing.T) {
	expected := []string{"(", "if", "(", ">", "2", "3", ")", ")"}
	actual := Tokenize("( if ( > 2 3 ) )")
	assert.Equal(t, expected, actual, fmt.Sprintf("Error, expected %s, actual %s", expected, actual))
}

func TestParseFile(t *testing.T) {
	expected := "(if (> 2 3) (print \"hello\") (if (< 4 7) ( print \"world\" ) ( print \"hello\" ) ) )"
	bytes, err := ioutil.ReadFile("../lisp.txt")
	if err != nil {
		log.Fatal(err)
	}
	actual := ParseFile(bytes)
	assert.Equal(t, expected, actual, fmt.Sprintf("Error, expected %s, actual %s", expected, actual))
}
