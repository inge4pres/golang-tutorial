package subset

import (
	"testing"
)

func TestSubsetSolver(t *testing.T) {
	target := 6
	current := []int{1, -2, 3, 8, 0, 3}
	SubsetSolver(current, target)
}
