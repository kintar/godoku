package godoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell_SetCandidateValues(t *testing.T) {
	c := new(Cell)
	c.SetCandidateValues([]int{1, 2, 3})
	assert.Equal(t, 7, int(c.Candidates))
	assert.Equal(t, []int{1, 2, 3}, c.GetCandidateValues())
}
