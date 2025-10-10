package parser

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestParseMovies_Table(t *testing.T) {
	tests := []struct {
		name, html string
		want       []Movie
	}{
		{"parse any movies",
			`
				<div class="session slick-slide slick-active">
				<a href="/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/">
					<div class="title">Ужастики. Ожившие рисунки </div>
					<div class="head">
					<div class="photo">
						<img src="https://cdn.mirage.ru/images/film/7000/small/p7003.jpg" alt="">
					</div>
					<div class="content">
						<div class="genre">комедия, фэнтези, семейный</div>
						<div class="labels"><span class="blue">Зал №2</span></div>
						<div class="bot">
						<div class="info-line">
							<div class="time">13:10</div>
							<div class="border">
							<div class="age">12+</div>
							<div class="format">2D</div>
							</div>
						</div>
						</div>
					</div>
					</div>
					<div class="place-price">
					<div class="item">500₽</div>
					</div>
				</a>
				</div>
				<div class="session slick-slide slick-active">
				<a href="/ticket/236100e0-456b-45b5-a719-6db9e44ffcc8/">
					<div class="title">Комедия. Бегущие за счастьем </div>
					<div class="head">
					<div class="photo">
						<img src="https://cdn.mirage.ru/images/film/7000/small/p7004.jpg" alt="">
					</div>
					<div class="content">
						<div class="genre">комедия, мелодрама, фантастика</div>
						<div class="labels"><span class="blue">Зал №3</span></div>
						<div class="bot">
						<div class="info-line">
							<div class="time">18:25</div>
							<div class="border">
							<div class="age">6+</div>
							<div class="format">3D</div>
							</div>
						</div>
						</div>
					</div>
					</div>
					<div class="place-price">
					<div class="item">550₽</div>
					</div>
				</a>
				</div>
				`,

			[]Movie{{
				Title:  "Ужастики. Ожившие рисунки",
				Time:   "13:10",
				Hall:   "Зал №2",
				Genre:  "комедия, фэнтези, семейный",
				Price:  "500₽",
				Age:    "12+",
				URL:    "/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/",
				Format: "2D",
			}, {
				Title:  "Комедия. Бегущие за счастьем",
				Time:   "18:25",
				Hall:   "Зал №3",
				Genre:  "комедия, мелодрама, фантастика",
				Price:  "550₽",
				Age:    "6+",
				URL:    "/ticket/236100e0-456b-45b5-a719-6db9e44ffcc8/",
				Format: "3D",
			}},
		}, {"parse one movies",
			`
				<div class="session slick-slide slick-active">
				<a href="/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/">
					<div class="title">Ужастики. Ожившие рисунки </div>
					<div class="head">
					<div class="photo">
						<img src="https://cdn.mirage.ru/images/film/7000/small/p7003.jpg" alt="">
					</div>
					<div class="content">
						<div class="genre">комедия, фэнтези, семейный</div>
						<div class="labels"><span class="blue">Зал №2</span></div>
						<div class="bot">
						<div class="info-line">
							<div class="time">13:10</div>
							<div class="border">
							<div class="age">12+</div>
							<div class="format">2D</div>
							</div>
						</div>
						</div>
					</div>
					</div>
					<div class="place-price">
					<div class="item">500₽</div>
					</div>
				</a>
				</div>
				`,
			[]Movie{{
				Title:  "Ужастики. Ожившие рисунки",
				Time:   "13:10",
				Hall:   "Зал №2",
				Genre:  "комедия, фэнтези, семейный",
				Price:  "500₽",
				Age:    "12+",
				URL:    "/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/",
				Format: "2D",
			}},
		},
		{"empty data",
			" ",
			[]Movie{},
		},
		{"incomplete date",
			`
				<div class="session slick-slide slick-active">
				<a href="/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/">
					<div class="title">Ужастики. Ожившие рисунки </div>
					<div class="head">
					<div class="photo">
						<img src="https://cdn.mirage.ru/images/film/7000/small/p7003.jpg" alt="">
					</div>
					<div class="content">
						<div class="genre">комедия, фэнтези, семейный</div>
						<div class="labels"><span class="blue">Зал №2</span></div>
						<div class="bot">
						<div class="info-line">
							<div class="time"></div>
							<div class="border">
							<div class="age">12+</div>
							<div class="format">2D</div>
							</div>
						</div>
						</div>
					</div>
					</div>
					<div class="place-price">
					<div class="item">500₽</div>
					</div>
				</a>
				</div>
				`,
			[]Movie{{
				Title:  "Ужастики. Ожившие рисунки",
				Time:   "",
				Hall:   "Зал №2",
				Genre:  "комедия, фэнтези, семейный",
				Price:  "500₽",
				Age:    "12+",
				URL:    "/ticket/236100e0-456b-45b5-a719-6db9e44ffcc9/",
				Format: "2D",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movies, _ := ParseMovies(strings.NewReader(tt.html))
			got, _ := json.Marshal(movies)
			want, _ := json.Marshal(tt.want)
			if !bytes.Equal(got, want) {
				t.Fatalf("json mismatch:\n got=%s\nwant=%s", got, want)
			}
		})
	}

}

func TestNormalizeText_Table(t *testing.T) {
	tests := []struct {
		name, text, want string
	}{
		{"trim multiple spaces", "hi,  people", "hi, people"},
		{"normalize case", "HeLLo, WoRld hIi WoRld", "Hello, World hii World"},
		{"html entities", "ok&bnsp", "ok"},
		{"removing line breaks", "hi, people\n\t\r\n", "hi, people"},
		{"trim spaces start and end", " hi, people ", "hi, people"},
	}

	for _, tt := range tests {
		ttt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeText(tt.text)
			if got != ttt.want {
				t.Fatalf("got %q want %q", got, tt.want)
			}
		})
	}
}
