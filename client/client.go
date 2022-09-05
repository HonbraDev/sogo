package client

import (
	"github.com/HonbraDev/sogo/base"
)

type Client struct {
	base.BaseClient
}

func NewClient(usr, pwd string) *Client {
	base := base.NewBaseClient(usr, pwd)
	return &Client{*base}
}
