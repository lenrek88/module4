package parser

import (
	"io"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

type Movie struct {
	Title  string
	Time   string
	Hall   string
	Genre  string
	Price  string
	Age    string
	URL    string
	Format string
}

func ParseMovies(r io.Reader) ([]Movie, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	movies := make([]Movie, 0)

	doc.Find(".session").Each(func(i int, card *goquery.Selection) {
		url, _ := card.Find("a").Attr("href")
		movie := Movie{
			Title:  NormalizeText(card.Find(".title").Text()),
			Genre:  NormalizeText(card.Find(".genre").Text()),
			Hall:   NormalizeText(card.Find(".labels span").Text()),
			Time:   NormalizeText(card.Find(".time").Text()),
			Age:    NormalizeText(card.Find(".age").Text()),
			Format: NormalizeText(card.Find(".format").Text()),
			Price:  NormalizeText(card.Find(".place-price .item").Text()),
			URL:    url,
		}
		movies = append(movies, movie)
	})
	return movies, nil
}

func NormalizeText(text string) string {
	if len(text) == 0 {
		return text
	}

	textFields := strings.Fields(text)
	for i, word := range textFields {
		runes := []rune(word)
		for j := 0; j < len(runes); j++ {
			if j != 0 {
				runes[j] = unicode.ToLower(runes[j])
			}
		}
		textFields[i] = string(runes)
	}

	text = strings.Join(textFields, " ")

	text = strings.Join(strings.Fields(text), " ")
	text = strings.ReplaceAll(text, "&bnsp", "")
	return text
}
