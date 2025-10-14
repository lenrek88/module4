package model

import (
	"errors"
	"strings"
	"time"
)

type Genre struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

func (g *Genre) Validate() error {
	if strings.TrimSpace(g.Title) == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
