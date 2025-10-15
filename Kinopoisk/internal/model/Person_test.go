package model

import (
	"testing"
	"time"
)

func TestDirectorValidate(t *testing.T) {
	t.Run("FirstName Empty", func(t *testing.T) {
		person := Person{
			ID:        12,
			FirstName: "",
			LastName:  "Fader",
			BirthDay:  time.Date(1975, time.October, 17, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
		}
		got := person.Validate()
		if got == nil {
			t.Fatalf("got %q want error FirstName do not Empty", got)
		}
	})

	t.Run("LastName Empty", func(t *testing.T) {
		person := Person{
			ID:        12,
			FirstName: "Alex",
			LastName:  "",
			BirthDay:  time.Date(1975, time.October, 17, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
		}
		got := person.Validate()
		if got == nil {
			t.Fatalf("got %q want error LastName do not Empty", got)
		}
	})

	t.Run("The birthday must be before the present time", func(t *testing.T) {
		person := Person{
			ID:        12,
			FirstName: "Alex",
			LastName:  "Fader",
			BirthDay:  time.Date(4343, time.October, 17, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
		}
		got := person.Validate()
		if got == nil {
			t.Fatalf("got %q want error LastName do not Empty", got)
		}
	})

	t.Run("Valid person", func(t *testing.T) {
		person := Person{
			ID:        12,
			FirstName: "Alex",
			LastName:  "Fader",
			BirthDay:  time.Date(1975, time.October, 17, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, time.January, 21, 10, 30, 0, 0, time.UTC),
		}
		got := person.Validate()
		if got != nil {
			t.Fatalf("got %q want nil", got)
		}
	})
}
