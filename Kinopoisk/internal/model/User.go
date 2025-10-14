package model

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("title cannot be empty")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email")

	}
	return nil
}
