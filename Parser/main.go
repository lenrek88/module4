package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"lenrek88/parser"
	"net/http"
)

func main() {
	body, err := FetchSchedule()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	movies, err := parser.ParseMovies(body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	PrintMovies(movies)
}

func PrintMovies(movies []parser.Movie) {
	for _, e := range movies {
		fmt.Printf("%s\n", e.Title)
		fmt.Printf("Жанр:	%s\n", e.Genre)
		fmt.Printf("Зал:	%s\n", e.Hall)
		fmt.Printf("Время:	%s\n", e.Time)
		fmt.Printf("Возраст:	%s\n", e.Age)
		fmt.Printf("Формат:	%s\n", e.Format)
		fmt.Printf("Цена:	%s\n", e.Price)
		fmt.Printf("URL:	%s\n", e.URL)
		fmt.Printf("\n")
	}
}

func FetchSchedule() (io.Reader, error) {
	url := "https://www.mirage.ru/msk/schedule/"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bodyBytes), nil
}
