package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {
	tCases := []struct {
		x      bool
		y      bool
		result bool
	}{
		{
			x:      true,
			y:      true,
			result: false,
		},
		{
			x:      true,
			y:      false,
			result: true,
		},
		{
			x:      false,
			y:      true,
			result: true,
		},
		{
			x:      false,
			y:      false,
			result: false,
		},
	}

	for _, c := range tCases {
		t.Run("test xor", func(t *testing.T) {
			assert.Equal(t, c.result, xor(c.x, c.y))
		})
	}
}
