package repository

import (
	"context"

	"github.com/igefined/nftique/internal/domain"
	"github.com/igefined/nftique/pkg/db/postgresql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type UserRepository struct {
	qb *postgresql.QBuilder
}

func NewUserRepository(qb *postgresql.QBuilder) *UserRepository {
	return &UserRepository{qb: qb}
}

func (r *UserRepository) GetByWeb3Address(ctx context.Context, web3Address string) (usr *domain.User, err error) {
	q := `select id, web3_address, first_name, last_name, username, created_at, updated_at, deactivated_at 
from users where web3_address = $1`
	err = pgxscan.Get(ctx, r.qb.Querier(), &usr, q, web3Address)
	return
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (err error) {
	q := `insert into users(id, web3_address, first_name, last_name, username) values ($1, $2, $3, $4, $5)`
	_, err = r.qb.Querier().Exec(ctx, q, user.UID, user.Web3Address, user.FirstName, user.LastName, user.Username)
	return
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (usr *domain.User, err error) {
	q := `select id, web3_address, first_name, last_name, username, created_at, updated_at, deactivated_at 
from users where username = $1`
	err = pgxscan.Get(ctx, r.qb.Querier(), &usr, q, username)
	return
}

func (r *UserRepository) DeactivateUser(ctx context.Context, uid uuid.UUID) (err error) {
	q := `delete from users where id = $1`
	_, err = r.qb.Querier().Exec(ctx, q, uid)
	return
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	uid uuid.UUID,
	user *domain.User,
) (usr *domain.User, err error) {
	return user, nil
}
