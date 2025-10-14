package model

import (
	"errors"
	"strings"
	"time"
)

type Validator interface {
	Validate() error
}

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	BirthDay  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Person) Validate() error {
	if strings.TrimSpace(p.FirstName) == "" {
		return errors.New("FirstName cannot be empty")
	}
	if strings.TrimSpace(p.LastName) == "" {
		return errors.New("LastName cannot be empty")
	}
	if p.BirthDay.After(time.Now()) {
		return errors.New("birth day cannot be in the future")
	}

	return nil
}
