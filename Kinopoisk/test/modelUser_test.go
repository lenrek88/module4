package test

import (
	"lenrek88/internal/model"
	"testing"
	"time"
)

func TestUserValidate(t *testing.T) {
	t.Run("Title Empty", func(t *testing.T) {
		user := model.User{
			ID:        12,
			Name:      "",
			Email:     "mail@yandex.ru",
			IsAdmin:   false,
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := user.Validate()
		if got == nil {
			t.Fatalf("got %q want error Title do not Empty", got)
		}
	})
	t.Run("Parse Email", func(t *testing.T) {
		user := model.User{
			ID:        12,
			Name:      "Alex",
			Email:     "mail.ru",
			IsAdmin:   false,
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := user.Validate()
		if got == nil {
			t.Fatalf("got %q want error for parse email", got)
		}
	})

	t.Run("Valid genre", func(t *testing.T) {
		user := model.User{
			ID:        12,
			Name:      "Alex",
			Email:     "kernel@mail.ru",
			IsAdmin:   false,
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := user.Validate()
		if got != nil {
			t.Fatalf("got %q want nil", got)
		}
	})

}
