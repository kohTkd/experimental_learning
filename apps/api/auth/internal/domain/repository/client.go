package repository

import "github.com/kohTkd/experimental_learning/ent"

type (
	Client struct {
		Client *ent.Client
	}
	Store interface {
		Transaction(fn func(cl *Client) error) error
		ReadOnly(fn func(cl *Client) error) error
		Command(fn func(cl *Client) error) error
	}
)

func NewClient(client *ent.Client) *Client {
	return &Client{Client: client}
}
