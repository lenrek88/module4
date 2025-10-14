package model

import (
	"errors"
	"strings"
	"time"
)

type UserID int64

type UserRate struct {
	UserID UserID
	Value  float64
}

type Film struct {
	ID         int64
	Title      string
	Year       int
	Duration   int
	DirectorID int64
	GenreIDs   []int64
	ActorIDs   []int64
	Rates      []UserRate
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (f *Film) Validate() error {
	if strings.TrimSpace(f.Title) == "" {
		return errors.New("title cannot be empty")
	}
	if f.Year < 1895 || f.Year > 2025 {
		return errors.New("year must be between 1895 and 2025")
	}
	if f.Duration <= 0 {
		return errors.New("duration must be positive")
	}
	return nil
}
