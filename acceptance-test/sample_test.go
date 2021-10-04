// +build integration

package acceptancetest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSystemCanReportErrors(t *testing.T) {
// 	result := 11
//
// 	assert.Equal(t, 10, result)
// }

func TestSystemWorks(t *testing.T) {
	result := 11

	assert.Equal(t, 11, result)
}
