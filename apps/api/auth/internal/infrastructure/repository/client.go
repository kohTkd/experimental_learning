package repository

import (
	"fmt"

	"github.com/kohTkd/experimental_learning/ent"
	"github.com/kohTkd/experimental_learning/internal/config"

	"entgo.io/ent/dialect"
)

type Client struct {
	Reader *ent.Client
	Writer *ent.Client
}

func NewClient() (*Client, error) {
	c := &Client{}

	var err error

	c.Reader, err = ent.Open(dialect.MySQL, config.Conf.Database.Reader.Dsn())
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	c.Writer, err = ent.Open(dialect.MySQL, config.Conf.Database.Writer.Dsn())
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	return c, nil
}

func (c *Client) Close() {
	c.Reader.Close()
	c.Writer.Close()
}
