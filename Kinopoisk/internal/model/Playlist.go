package model

import (
	"errors"
	"strings"
	"time"
)

type Playlist struct {
	ID        int64
	Title     string
	OwnerID   string
	FilmIds   []int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Playlist) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
