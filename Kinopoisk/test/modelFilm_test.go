package test

import (
	"encoding/json"
	"fmt"
	"lenrek88/internal/model"
	"os"
	"path/filepath"
	"testing"
)

func loadFilmFixture(filename string) model.Film {
	data, err := os.ReadFile(filepath.Join("testdata", filename))
	if err != nil {
		panic(fmt.Sprintf("failed to load fixture: %v", err))
	}

	var film model.Film
	if err := json.Unmarshal(data, &film); err != nil {
		panic(fmt.Sprintf("failed to parse fixture: %v", err))
	}

	return film
}

/// Фикстуры разобрался, а моки что добавить? Плюс юзеру добавить

func validFilmFixture() model.Film {
	return loadFilmFixture("filmFixture_valide.json")
}

func emptyTitleFilmFixture() model.Film {
	f := validFilmFixture()
	f.Title = ""
	return f
}

func TestValidateFilm_Table(t *testing.T) {
	tests := []struct {
		name        string
		fixture     func() model.Film
		expectError bool
	}{
		{name: "valid film",
			fixture:     validFilmFixture,
			expectError: false,
		},
		{name: "empty Title",
			fixture:     emptyTitleFilmFixture,
			expectError: true,
		},
	}
	for _, tt := range tests {
		ttt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			film := ttt.fixture()
			err := film.Validate()
			if ttt.expectError && err == nil {
				t.Error("expected error, got nil")
			}
			if !ttt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// func TestValidateFilm_Table(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		film model.Film
// 	}{
// 		{"empty Title", model.Film{
// 			ID:         1,
// 			Title:      "",
// 			Year:       2025,
// 			Duration:   138,
// 			DirectorID: 2352,
// 			GenreIDs:   []int64{1, 3, 4},
// 			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
// 			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
// 			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
// 			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
// 		}},
// 		{"year must be between 1895-2025", model.Film{
// 			ID:         1,
// 			Title:      "Август",
// 			Year:       1894,
// 			Duration:   138,
// 			DirectorID: 2352,
// 			GenreIDs:   []int64{1, 3, 4},
// 			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
// 			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
// 			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
// 			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
// 		}},
// 		{"year must be between 1895-2025", model.Film{
// 			ID:         1,
// 			Title:      "Август",
// 			Year:       2026,
// 			Duration:   138,
// 			DirectorID: 2352,
// 			GenreIDs:   []int64{1, 3, 4},
// 			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
// 			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
// 			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
// 			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
// 		}},
// 		{"duration = 0", model.Film{
// 			ID:         1,
// 			Title:      "Август",
// 			Year:       2025,
// 			Duration:   0,
// 			DirectorID: 2352,
// 			GenreIDs:   []int64{1, 3, 4},
// 			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
// 			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
// 			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
// 			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
// 		}},
// 		{"duration < 0", model.Film{
// 			ID:         1,
// 			Title:      "Август",
// 			Year:       2025,
// 			Duration:   -1,
// 			DirectorID: 2352,
// 			GenreIDs:   []int64{1, 3, 4},
// 			ActorIDs:   []int64{1, 3, 5, 22, 53, 45, 11, 34, 99, 21},
// 			Rates:      []model.UserRate{{UserID: 25, Value: 5.4}, {UserID: 12, Value: 9.0}},
// 			CreatedAt:  time.Date(2025, time.September, 25, 0, 0, 0, 0, time.UTC),
// 			UpdatedAt:  time.Date(2025, time.September, 26, 0, 0, 0, 0, time.UTC),
// 		}},
// 	}

// 	for _, tt := range tests {
// 		ttt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			got := ttt.film.Validate()
// 			if got == nil {
// 				t.Fatalf("got %q want error %q", got, tt.name)
// 			}
// 		})
// 	}
// }
