// +build integration

package acceptancetest

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemWorks(t *testing.T) {
	output, err := exec.Command("../scheme", "acceptance-test/data/simple_expression.sm").Output()
	assert.NoError(t, err)
	assert.Equal(t, "4\n", string(output))

}
