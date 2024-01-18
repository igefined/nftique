package service

import "github.com/ethereum/go-ethereum/ethclient"

type Contractor struct {
	client *ethclient.Client
}

func NewContractor(client *ethclient.Client) *Contractor {
	return &Contractor{client: client}
}
