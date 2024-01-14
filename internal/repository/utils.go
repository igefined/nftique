package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsDuplicate(err error) bool {
	var pgErr *pgconn.PgError
	if ok := errors.As(err, &pgErr); ok {
		if pgErr == nil {
			return false
		}

		if pgErr.Code == "23505" {
			return true
		}
	}

	return false
}
