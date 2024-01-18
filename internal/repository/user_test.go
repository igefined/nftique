package repository

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/igefined/nftique/internal/domain"
)

func (s *Suite) TestUserRepository_Get() {
	s.Run("success, exist", func() {
		fUsr, err := s.createUser(s.ctx)
		s.Require().NoError(err)

		usr, err := s.repository.Get(s.ctx, fUsr.UID)
		s.Require().NoError(err)
		s.Require().Equal(usr.UID, fUsr.UID)
	})

	s.Run("success, not exist", func() {
		uid := uuid.New()
		usr, err := s.repository.Get(s.ctx, uid)
		s.Require().Nil(usr)
		s.Require().Error(err, domain.ErrEntityNotFound)
	})

	s.clearDB()
}

func (s *Suite) TestUserRepository_GetByWeb3Address() {
	s.Run("success, exist", func() {
		fUsr, err := s.createUser(s.ctx)
		s.Require().NoError(err)

		usr, err := s.repository.GetByWeb3Address(s.ctx, fUsr.Web3Address)
		s.Require().NoError(err)
		s.Require().Equal(usr.UID, fUsr.UID)
	})

	s.Run("success, not exist", func() {
		web3Address := "0x3E1D0cEd18A4454BA390b8F540682c718748b0e6"
		usr, err := s.repository.GetByWeb3Address(s.ctx, web3Address)
		s.Require().Nil(usr)
		s.Require().Error(err, domain.ErrEntityNotFound)
	})

	s.clearDB()
}

func (s *Suite) TestUserRepository_GetByUsername() {
	s.Run("success, exist", func() {
		fUsr, err := s.createUser(s.ctx)
		s.Require().NoError(err)

		usr, err := s.repository.GetByUsername(s.ctx, fUsr.Username)
		s.Require().NoError(err)
		s.Require().Equal(usr.UID, fUsr.UID)
	})

	s.Run("success, not exist", func() {
		username := "invalid_username"
		usr, err := s.repository.GetByUsername(s.ctx, username)
		s.Require().Nil(usr)
		s.Require().Error(err, domain.ErrEntityNotFound)
	})

	s.clearDB()
}

func (s *Suite) TestUserRepository_CreateUser() {
	s.Run("success", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		usr, err := s.repository.GetByUsername(s.ctx, fUsr.Username)
		s.Require().NoError(err)
		s.Require().Equal(usr.UID, fUsr.UID)
	})

	s.Run("duplicate web3_address", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := fakeUser()
		fUsr2.Web3Address = fUsr.Web3Address
		err = s.repository.CreateUser(s.ctx, fUsr2)
		s.Require().Error(err, domain.ErrEntityDuplicate)
	})

	s.clearDB()
}

func (s *Suite) TestUserRepository_DeactivateUser() {
	s.Run("success", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		err = s.repository.DeactivateUser(s.ctx, fUsr.UID)
		s.Require().NoError(err)

		usr, err := s.repository.GetByUsername(s.ctx, fUsr.Username)
		s.Require().NoError(err)
		s.Require().Equal(usr.UID, fUsr.UID)
		s.Require().NotNil(usr.DeactivatedAt)
	})

	s.clearDB()
}

func (s *Suite) TestUserRepository_UpdateUser() {
	s.Run("success, all params", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := fakeUser()
		fUsr2.UID = fUsr.UID

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().NoError(err)
		s.Require().Equal(usr.LastName, fUsr2.LastName)
	})

	s.Run("success, only last name", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().NoError(err)
		s.Require().Equal(usr.LastName, fUsr2.LastName)
	})

	s.Run("success, only username", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			Username: randomdata.SillyName(),
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().NoError(err)
		s.Require().Equal(usr.Username, fUsr2.Username)
	})

	s.Run("success, only first name", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := &domain.User{
			UID:       fUsr.UID,
			FirstName: randomdata.FirstName(randomdata.RandomGender),
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().NoError(err)
		s.Require().Equal(usr.FirstName, fUsr2.FirstName)
	})

	s.Run("success, username and last name", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
			Username: randomdata.SillyName(),
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().NoError(err)
		s.Require().Equal(usr.LastName, fUsr2.LastName)
		s.Require().Equal(usr.Username, fUsr2.Username)
	})

	s.Run("err, user does not exist", func() {
		fUsr := fakeUser()

		fUsr2 := &domain.User{
			UID:      fUsr.UID,
			LastName: randomdata.LastName(),
			Username: randomdata.SillyName(),
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr.UID, fUsr2)
		s.Require().Error(err, domain.ErrEntityNotFound)
		s.Require().Nil(usr)
	})

	s.Run("err, username already exist", func() {
		fUsr := fakeUser()
		err := s.repository.CreateUser(s.ctx, fUsr)
		s.Require().NoError(err)

		fUsr2 := fakeUser()
		err = s.repository.CreateUser(s.ctx, fUsr2)
		s.Require().NoError(err)

		fUsr3 := &domain.User{
			UID:      fUsr2.UID,
			Username: fUsr.Username,
		}

		usr, err := s.repository.UpdateUser(s.ctx, fUsr2.UID, fUsr3)
		s.Require().Nil(usr)
		s.Require().Error(err, domain.ErrEntityDuplicate)
	})

	s.clearDB()
}
