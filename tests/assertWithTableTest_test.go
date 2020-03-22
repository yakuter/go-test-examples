package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(test.expected, Calculate(test.input))
	}
}
