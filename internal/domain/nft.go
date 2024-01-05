package domain

import "time"

type NFT struct {
	Creator   string    `json:"creator"`
	Owner     string    `json:"owner"`
	Img       string    `json:"img"`
	CreatedAt time.Time `json:"created_at"`
}
