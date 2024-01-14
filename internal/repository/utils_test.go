package repository

import (
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestIsDuplicate(t *testing.T) {
	testCases := []struct {
		pgErr    *pgconn.PgError
		expected bool
	}{
		{
			pgErr: &pgconn.PgError{
				Code: "23505",
			},
			expected: true,
		},
		{
			pgErr: &pgconn.PgError{
				Code: "23504",
			},
			expected: false,
		},
		{
			pgErr:    &pgconn.PgError{},
			expected: false,
		},
		{
			pgErr:    nil,
			expected: false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, IsDuplicate(tc.pgErr))
	}
}
