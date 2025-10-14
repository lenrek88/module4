package test

import (
	"lenrek88/internal/model"
	"testing"
	"time"
)

func TestValidateFilm(t *testing.T) {
	film := model.Film{
		ID:         1,
		Title:      "Август",
		Year:       2025,
		Duration:   138,
		DirectorID: 2352,
		GenreIDs:   []int64{1, 3, 4},
		ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
		Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
		CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
	}
	err := film.Validate()
	if err != nil {
		t.Fatalf("got err %s, want nil", err)

	}
}

func TestValidateFilm_Table(t *testing.T) {
	tests := []struct {
		name string
		film model.Film
	}{
		{"empty Title", model.Film{
			ID:         1,
			Title:      "",
			Year:       2025,
			Duration:   138,
			DirectorID: 2352,
			GenreIDs:   []int64{1, 3, 4},
			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
		}},
		{"year must be between 1895-2025", model.Film{
			ID:         1,
			Title:      "Август",
			Year:       1894,
			Duration:   138,
			DirectorID: 2352,
			GenreIDs:   []int64{1, 3, 4},
			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
		}},
		{"year must be between 1895-2025", model.Film{
			ID:         1,
			Title:      "Август",
			Year:       2026,
			Duration:   138,
			DirectorID: 2352,
			GenreIDs:   []int64{1, 3, 4},
			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
		}},
		{"duration = 0", model.Film{
			ID:         1,
			Title:      "Август",
			Year:       2025,
			Duration:   0,
			DirectorID: 2352,
			GenreIDs:   []int64{1, 3, 4},
			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
		}},
		{"duration < 0", model.Film{
			ID:         1,
			Title:      "Август",
			Year:       2025,
			Duration:   -1,
			DirectorID: 2352,
			GenreIDs:   []int64{1, 3, 4},
			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
		}},
	}

	for _, tt := range tests {
		ttt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ttt.film.Validate()
			if got == nil {
				t.Fatalf("got %q want error %q", got, tt.name)
			}
		})
	}
}
