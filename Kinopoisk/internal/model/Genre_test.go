package model

import (
	"testing"
	"time"
)

func TestGendreValidate(t *testing.T) {
	t.Run("Title Empty", func(t *testing.T) {
		genre := Genre{
			ID:        12,
			Title:     "",
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := genre.Validate()
		if got == nil {
			t.Fatalf("got %q want error Title do not Empty", got)
		}
	})

	t.Run("Valid genre", func(t *testing.T) {
		genre := Genre{
			ID:        12,
			Title:     "Боевик",
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := genre.Validate()
		if got != nil {
			t.Fatalf("got %q want nil", got)
		}
	})

}
