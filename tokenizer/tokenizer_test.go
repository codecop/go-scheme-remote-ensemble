/* Tokenizer
- input:
    - string
    - cleaned
- output:
    - list of tokens
        - what type is it? (e.g. slice, linked list, interface, struct)
        - brackets, identifiers (func names, variables), values (number, "string", boolean #t#f, other literals)
        - ignore symbols
        - (extend later if necessary)
    - token (struct, interface)
        - value (e.g. identifier -> string, number -> number, bracket -> nil)
        - type
*/
package tokenizer_test

import (
	"testing"
	"tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStringTokenizesIntoEmpty(t *testing.T) {
	cleanedString := ""
	tokens := tokenizer.Scan(cleanedString)
	assert.Equal(t, make([]tokenizer.Token, 0), tokens)
}
