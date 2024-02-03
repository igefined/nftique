package service

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/igefined/nftique/internal/domain"
	"github.com/igefined/nftique/pkg/config"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate mockgen -source=auth.go -package=mocks -destination=./mocks/mock_user.go UserRepository

type UserRepository interface {
	GetByWeb3Address(ctx context.Context, web3Address string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, uid uuid.UUID, user *domain.User) (*domain.User, error)
	DeactivateUser(ctx context.Context, uid uuid.UUID) error
}

type NFTService struct {
	cfg    *config.ETHCfg
	logger *zap.Logger

	client             *ethclient.Client
	operatorAddress    common.Address
	nftMarketplaceAddr common.Address
}

func NewNFTService(
	cfg *config.ETHCfg,
	logger *zap.Logger,
) (*NFTService, error) {
	client, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("error create eth client")
	}

	privKey, err := getPrivateKey(cfg)
	if err != nil {
		return nil, fmt.Errorf("generate private key error, err: %s", err.Error())
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("get public key error")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !bytes.Equal(common.HexToAddress(cfg.OperatorAddr).Bytes(), fromAddress.Bytes()) {
		return nil, fmt.Errorf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.OperatorAddr, fromAddress,
		)
	}

	return &NFTService{
		cfg:                cfg,
		logger:             logger,
		client:             client,
		operatorAddress:    common.HexToAddress(cfg.OperatorAddr),
		nftMarketplaceAddr: common.HexToAddress(cfg.ContractAddress),
	}, nil
}

func getPrivateKey(config *config.ETHCfg) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func (s *NFTService) ListAllAvailable(_ context.Context) ([]*domain.NFT, error) {
	size := rand.Intn(10)
	mockNFTs := make([]*domain.NFT, size)

	for i := range mockNFTs {
		mockNFTs[i] = &domain.NFT{
			Creator:   "0x3E1D0cEd18A4454BA390b8F540682c718748b0e5",
			Owner:     "0x3E1D0cEd18A4454BA390b8F540682c718748b0e5",
			CreatedAt: time.Now().Add(-time.Minute * 2),
			Img:       "https://nftique.s3.amazonaws.com/63f902d79a33f77a446ce12c_qIsIXihKZeVDop6AZWt1j6gxOnYZ_oGfr09PzlJDL_H4YWasvDrNuTXK8Qrmh0oL6ppWI3RaGU5vMif2gNwO38UdWWei4eZCNhbfdrVlm5qHV3zVYIk6qtBuFvdkoo0HexhmSmvn.jpeg", //nolint:lll
		}
	}

	return mockNFTs, nil
}
