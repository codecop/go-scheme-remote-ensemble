// +build integration

package acceptancetest

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemWorks(t *testing.T) {
	output, err := exec.Command("../scheme", "acceptance-test/data/simple_expression.scm").Output()
	assert.Equal(t, "4\n", string(output))
	assert.NoError(t, err)
}

// TODO rename the file to a more reasonable name
