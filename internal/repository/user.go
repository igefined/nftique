package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/igefined/nftique/internal/domain"
	"github.com/igefined/nftique/pkg/db/postgresql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	qb *postgresql.QBuilder
}

func NewUserRepository(qb *postgresql.QBuilder) *UserRepository {
	return &UserRepository{qb: qb}
}

func (r *UserRepository) Get(ctx context.Context, uid uuid.UUID) (*domain.User, error) {
	var usr domain.User
	q := `select * from users where id = $1`
	if err := pgxscan.Get(ctx, r.qb.Querier(), &usr, q, uid); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrEntityNotFound
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepository) GetByWeb3Address(ctx context.Context, web3Address string) (*domain.User, error) {
	var usr domain.User
	q := `select * from users where web3_address = $1`
	if err := pgxscan.Get(ctx, r.qb.Querier(), &usr, q, web3Address); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrEntityNotFound
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	q := `insert into users(id, web3_address, first_name, last_name, username) values ($1, $2, $3, $4, $5)`
	if _, err := r.qb.Querier().Exec(
		ctx,
		q,
		user.UID,
		user.Web3Address,
		user.FirstName,
		user.LastName,
		user.Username,
	); err != nil {
		if IsDuplicate(err) {
			return domain.ErrEntityDuplicate
		}

		return err
	}

	return nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var usr domain.User
	q := `select * from users where username = $1`
	if err := pgxscan.Get(ctx, r.qb.Querier(), &usr, q, username); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrEntityNotFound
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepository) DeactivateUser(ctx context.Context, uid uuid.UUID) (err error) {
	q := `update users set deactivated_at = now() where id = $1`
	_, err = r.qb.Querier().Exec(ctx, q, uid)
	return
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	uid uuid.UUID,
	user *domain.User,
) (usr *domain.User, err error) {
	usr, err = r.Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	query := "update users set "
	args := make([]any, 0)
	var iParam = 1

	if user.Username != "" {
		u, err := r.GetByUsername(ctx, user.Username)
		if u != nil && u.UID != uid {
			return nil, domain.ErrEntityDuplicate
		}

		if err != nil && !errors.Is(err, domain.ErrEntityNotFound) {
			return nil, err
		}

		query += fmt.Sprintf("username = $%d, ", iParam)
		args = append(args, user.Username)
		usr.Username = user.Username
		iParam++
	}

	if user.FirstName != "" {
		query += fmt.Sprintf("first_name = $%d, ", iParam)
		args = append(args, user.FirstName)
		usr.FirstName = user.FirstName
		iParam++
	}

	if user.LastName != "" {
		query += fmt.Sprintf("last_name = $%d, ", iParam)
		args = append(args, user.LastName)
		usr.LastName = user.LastName
		iParam++
	}

	query += "updated_at = now(), "
	query = query[:len(query)-2]

	query += fmt.Sprintf(" WHERE id = $%d", iParam)
	args = append(args, uid)

	_, err = r.qb.Querier().Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
