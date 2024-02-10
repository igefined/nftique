package domain

import "time"

type NFT struct {
	Creator   string    `json:"creator" validate:"eth_addr"`
	Owner     string    `json:"owner" validate:"eth_addr"`
	Img       string    `json:"img"`
	CreatedAt time.Time `json:"created_at"`
}

type NFTMedia struct {
	Date         string
	Url          string
	LastModified time.Time
}
