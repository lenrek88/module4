package parser

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Movie struct {
	Title string
	Time  string
	Hall  string
	Price string
}

func ParseMovies(r io.Reader) ([]Movie, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)
	doc.Find(".movie-card").Each(func(i int, card *goquery.Selection) {
		movie := Movie{
			Title: card.Find(".title").Text(),
			Time:  card.Find(".time").Text(),
			Hall:  card.Find(".hall").Text(),
			Price: card.Find(".price").Text(),
		}
		movies = append(movies, movie)
	})
	return movies, nil
}
