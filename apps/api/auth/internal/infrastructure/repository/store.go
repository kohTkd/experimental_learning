package repository

import (
	"context"
	"fmt"

	"github.com/kohTkd/experimental_learning/ent"
	"github.com/kohTkd/experimental_learning/internal/config"
	"github.com/kohTkd/experimental_learning/internal/domain/repository"

	"entgo.io/ent/dialect"
)

type Store struct {
	reader *repository.Client
	writer *repository.Client
}

func NewStore() (*Store, error) {
	s := &Store{}

	r, err := ent.Open(dialect.MySQL, config.Conf.Database.Reader.Dsn())
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	s.reader = repository.NewClient(r)

	w, err := ent.Open(dialect.MySQL, config.Conf.Database.Writer.Dsn())
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	s.writer = repository.NewClient(w)

	return s, nil
}

func (s *Store) Close() {
	s.reader.Client.Close()
	s.writer.Client.Close()
}

func (s *Store) Reader() *repository.Client {
	return s.reader
}

func (s *Store) Writer() *repository.Client {
	return s.writer
}

func (s *Store) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := s.writer.Client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	ctx = ent.NewTxContext(ctx, tx)
	ctx = ent.NewContext(ctx, tx.Client())

	if err := fn(ctx); err != nil {
		txErr := err
		if err := tx.Rollback(); err != nil {
			return err
		}
		return txErr
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Command(ctx context.Context, fn func(ctx context.Context) error) error {
	ctx = ent.NewContext(ctx, s.writer.Client)
	return fn(ctx)
}

func (s *Store) ReadOnly(ctx context.Context, fn func(ctx context.Context) error) error {
	ctx = ent.NewContext(ctx, s.reader.Client)
	return fn(ctx)
}
