package collatz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollatz(t *testing.T) {
	assert.Equal(t, 1, len(Collatz(1)))
	assert.Equal(t, 2, len(Collatz(2)))
	assert.Equal(t, 6, len(Collatz(5)))
	assert.Equal(t, 24, len(Collatz(995)))
	assert.Equal(t, 133, len(Collatz(1233)))
}
