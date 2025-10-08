package parser

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestParseOneMovies(t *testing.T) {

	html := `
    <div class="movie-card">
      <h3 class="title">Фильм А</h3>
      <span class="time">18:30</span>
      <span class="hall">Зал 2</span>
      <span class="price">350 ₽</span>
    </div>`

	movies, err := ParseMovies(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}
	m := []Movie{{Title: "Фильм А", Time: "18:30", Hall: "Зал 2", Price: "350 ₽"}}
	got, _ := json.Marshal(movies)
	want, _ := json.Marshal(m)

	if !bytes.Equal(got, want) {
		t.Fatalf("json mismatch:\n got=%s\nwant=%s", got, want)
	}
}

func TestParseAnyMovies(t *testing.T) {

	html := `
    <div class="movie-card">
      <h3 class="title">Фильм А</h3>
      <span class="time">18:30</span>
      <span class="hall">Зал 2</span>
      <span class="price">350 ₽</span>
    </div>
	<div class="movie-card">
      <h3 class="title">Фильм Б</h3>
      <span class="time">21:30</span>
      <span class="hall">Зал 3</span>
      <span class="price">420 ₽</span>
    </div>
	<div class="movie-card">
      <h3 class="title">Фильм В</h3>
      <span class="time">09:30</span>
      <span class="hall">Зал 1</span>
      <span class="price">550 ₽</span>
    </div>`

	movies, err := ParseMovies(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}
	m := []Movie{{Title: "Фильм А", Time: "18:30", Hall: "Зал 2", Price: "350 ₽"},
		{Title: "Фильм Б", Time: "21:30", Hall: "Зал 3", Price: "420 ₽"},
		{Title: "Фильм В", Time: "09:30", Hall: "Зал 1", Price: "550 ₽"}}
	got, _ := json.Marshal(movies)
	want, _ := json.Marshal(m)

	if !bytes.Equal(got, want) {
		t.Fatalf("json mismatch:\n got=%s\nwant=%s", got, want)
	}
}
