package test

import (
	"lenrek88/internal/model"
	"testing"
	"time"
)

func TestPlalistValidate(t *testing.T) {
	t.Run("Title Empty", func(t *testing.T) {
		playlist := model.Playlist{
			ID:        14,
			Title:     "",
			OwnerID:   "click click",
			FilmIds:   []int64{1, 2, 3},
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := playlist.Validate()
		if got == nil {
			t.Fatalf("got %q want error Title do not Empty", got)
		}
	})

	t.Run("Valid genre", func(t *testing.T) {
		playlist := model.Playlist{
			ID:        14,
			Title:     "My Playlist",
			OwnerID:   "click click",
			FilmIds:   []int64{1, 2, 3},
			CreatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, time.January, 20, 10, 30, 0, 0, time.UTC),
		}
		got := playlist.Validate()
		if got != nil {
			t.Fatalf("got %q want nil", got)
		}
	})

}
