package model

import (
	"encoding/json"
	"fmt"
	"lenrek88/internal/testutil"
	"testing"
)

/// Фикстуры разобрался, а моки что добавить? Плюс юзеру добавить

func validFilmFixture() Film {
	data := testutil.LoadFixture("model", "filmFixture_valide.json")

	var film Film
	if err := json.Unmarshal(data, &film); err != nil {
		panic(fmt.Sprintf("failed to parse fixture: %v", err))
	}
	return film
}

func emptyTitleFilmFixture() Film {
	f := validFilmFixture()
	f.Title = ""
	return f
}

func lowYearFilmFixture() Film {
	f := validFilmFixture()
	f.Year = 1894
	return f
}

func highYearFilmFixture() Film {
	f := validFilmFixture()
	f.Year = 2026
	return f
}

func zeroDurationFilmFixture() Film {
	f := validFilmFixture()
	f.Duration = 0
	return f
}

func negativeDurationFilmFixture() Film {
	f := validFilmFixture()
	f.Duration = -1
	return f
}

func TestValidateFilm_Table(t *testing.T) {
	tests := []struct {
		name        string
		fixture     func() Film
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
		{name: "low Year Film Fixture",
			fixture:     lowYearFilmFixture,
			expectError: true,
		},
		{name: "high Year Film Fixture",
			fixture:     highYearFilmFixture,
			expectError: true,
		},
		{name: "zero Duration Film Fixture",
			fixture:     zeroDurationFilmFixture,
			expectError: true,
		},
		{name: "negative Duration Film Fixture",
			fixture:     negativeDurationFilmFixture,
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
