package diffsq_test

import (
	"testing"

	"github.com/hsmtkk/crispy-memory/diffsq"
	"github.com/stretchr/testify/assert"
)

func TestDiffSquare(t *testing.T) {
	assert.Equal(t, 2640, diffsq.DiffSquare(10))
}

func TestSquareSum(t *testing.T) {
	assert.Equal(t, 3025, diffsq.SquareSum(10))
}

func TestSumSquare(t *testing.T) {
	assert.Equal(t, 385, diffsq.SumSquare(10))
}
