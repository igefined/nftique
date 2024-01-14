package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/igefined/nftique/internal/domain"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	repo := NewUserRepository(qb)

	t.Run("success, exist", func(t *testing.T) {
		ctx := context.Background()
		fUsr, err := createUser(ctx)
		assert.NoError(t, err)

		usr, err := repo.Get(ctx, fUsr.UID)
		assert.NoError(t, err)
		assert.Equal(t, usr.UID, fUsr.UID)
	})

	t.Run("success, not exist", func(t *testing.T) {
		ctx := context.Background()

		uid := uuid.New()
		usr, err := repo.Get(ctx, uid)
		assert.Nil(t, usr)
		assert.Error(t, err, domain.ErrEntityNotFound)
	})
}

func TestGetByWeb3Address(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	repo := NewUserRepository(qb)

	t.Run("success, exist", func(t *testing.T) {
		ctx := context.Background()
		fUsr, err := createUser(ctx)
		assert.NoError(t, err)

		usr, err := repo.GetByWeb3Address(ctx, fUsr.Web3Address)
		assert.NoError(t, err)
		assert.Equal(t, usr.UID, fUsr.UID)
	})

	t.Run("success, not exist", func(t *testing.T) {
		ctx := context.Background()

		web3Address := "0x3E1D0cEd18A4454BA390b8F540682c718748b0e6"
		usr, err := repo.GetByWeb3Address(ctx, web3Address)
		assert.Nil(t, usr)
		assert.Error(t, err, domain.ErrEntityNotFound)
	})
}

func TestGetByUsername(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	repo := NewUserRepository(qb)

	t.Run("success, exist", func(t *testing.T) {
		ctx := context.Background()
		fUsr, err := createUser(ctx)
		assert.NoError(t, err)

		usr, err := repo.GetByUsername(ctx, fUsr.Username)
		assert.NoError(t, err)
		assert.Equal(t, usr.UID, fUsr.UID)
	})

	t.Run("success, not exist", func(t *testing.T) {
		ctx := context.Background()

		username := "invalid_username"
		usr, err := repo.GetByUsername(ctx, username)
		assert.Nil(t, usr)
		assert.Error(t, err, domain.ErrEntityNotFound)
	})
}

func TestCreateUser(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	var repo = NewUserRepository(qb)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		usr, err := repo.GetByUsername(ctx, fUsr.Username)
		assert.NoError(t, err)
		assert.Equal(t, usr.UID, fUsr.UID)
	})

	t.Run("duplicate web3_address", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := fakeUser()
		fUsr2.Web3Address = fUsr.Web3Address
		err = repo.CreateUser(ctx, fUsr2)
		assert.Error(t, err, domain.ErrEntityDuplicate)
	})
}

func TestDeactivateUser(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	var repo = NewUserRepository(qb)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		err = repo.DeactivateUser(ctx, fUsr.UID)
		assert.NoError(t, err)

		usr, err := repo.GetByUsername(ctx, fUsr.Username)
		assert.NoError(t, err)
		assert.Equal(t, usr.UID, fUsr.UID)
		assert.NotNil(t, usr.DeactivatedAt)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Cleanup(func() {
		clearDB(t)
	})

	var repo = NewUserRepository(qb)

	t.Run("success, all params", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := fakeUser()
		fUsr2.UID = fUsr.UID

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.NoError(t, err)
		assert.Equal(t, usr.LastName, fUsr2.LastName)
	})

	t.Run("success, only last name", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
		}

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.NoError(t, err)
		assert.Equal(t, usr.LastName, fUsr2.LastName)
	})

	t.Run("success, only username", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			Username: randomdata.SillyName(),
		}

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.NoError(t, err)
		assert.Equal(t, usr.Username, fUsr2.Username)
	})

	t.Run("success, only first name", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := &domain.User{
			UID:       fUsr.UID,
			FirstName: randomdata.FirstName(randomdata.RandomGender),
		}

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.NoError(t, err)
		assert.Equal(t, usr.FirstName, fUsr2.FirstName)
	})

	t.Run("success, username and last name", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
			Username: randomdata.SillyName(),
		}

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.NoError(t, err)
		assert.Equal(t, usr.LastName, fUsr2.LastName)
		assert.Equal(t, usr.Username, fUsr2.Username)
	})

	t.Run("err, user does not exist", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
			Username: randomdata.SillyName(),
		}

		usr, err := repo.UpdateUser(ctx, fUsr.UID, fUsr2)
		assert.Error(t, err, domain.ErrEntityNotFound)
		assert.Nil(t, usr)
	})

	t.Run("err, username already exist", func(t *testing.T) {
		ctx := context.Background()
		fUsr := fakeUser()
		err := repo.CreateUser(ctx, fUsr)
		assert.NoError(t, err)

		fUsr2 := fakeUser()
		err = repo.CreateUser(ctx, fUsr2)
		assert.NoError(t, err)

		fUsr3 := &domain.User{
			UID:      fUsr2.UID,
			Username: fUsr.Username,
		}

		usr, err := repo.UpdateUser(ctx, fUsr2.UID, fUsr3)
		assert.Nil(t, usr)
		assert.Error(t, err, domain.ErrEntityDuplicate)
	})
}

func fakeUser() *domain.User {
	return &domain.User{
		UID:         uuid.New(),
		FirstName:   randomdata.FirstName(randomdata.RandomGender),
		LastName:    randomdata.LastName(),
		Username:    randomdata.PhoneNumber(),
		Web3Address: randomdata.Address(),
	}
}

func createUser(ctx context.Context) (*domain.User, error) {
	usr := fakeUser()
	q := `insert into users(id, web3_address, first_name, last_name, username) values ($1, $2, $3, $4, $5)`
	_, err := qb.Querier().Exec(ctx, q, usr.UID, usr.Web3Address, usr.FirstName, usr.LastName, usr.Username)
	if err != nil {
		return nil, errors.New("error inserted fake user")
	}

	return usr, err
}
