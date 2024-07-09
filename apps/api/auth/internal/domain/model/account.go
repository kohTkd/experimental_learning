package model

import (
	"github.com/kohTkd/experimental_learning/ent"
	"golang.org/x/crypto/bcrypt"
)

type Account = ent.Accounts

func NewAccount(email, password string) *Account {
	return &Account{
		Email:    email,
		Password: encryptPassword(password),
	}
}

func encryptPassword(password string) string {
	p := []byte(password)
	h, err := bcrypt.GenerateFromPassword(p, 12)
	if err != nil {
		panic(err)
	}
	return string(h)
}
