package repository

import (
	"fmt"

	"github.com/kohTkd/experimental_learning/ent"
	"github.com/kohTkd/experimental_learning/internal/config"

	"entgo.io/ent/dialect"

	"github.com/go-sql-driver/mysql"
)

type Client struct {
	Reader *ent.Client
	Writer *ent.Client
}

func NewClient() (*Client, error) {
	c := &Client{}

	var err error

	c.Reader, err = ent.Open(dialect.MySQL, dsn(config.Conf.Database.Reader))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	c.Writer, err = ent.Open(dialect.MySQL, dsn(config.Conf.Database.Writer))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	return c, nil
}

func (c *Client) Close() {
	c.Reader.Close()
	c.Writer.Close()
}

func dsn(conf config.DatabaseConfig) string {
	mc := mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  conf.Net,
		Addr:                 conf.Addr,
		DBName:               conf.Dbname,
		AllowNativePasswords: conf.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": conf.Params.ParseTime,
			"charset":   conf.Params.Charset,
			"loc":       conf.Params.Loc,
		},
	}

	return mc.FormatDSN()
}
