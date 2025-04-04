package main

import (
	"fmt"
	"os"

	"github.com/egordigitax/pneuma"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	p := pneuma.Init(os.Getenv("OPENAI_KEY"))

	type ClassifiedPoll struct {
		FriendshipLevel int    `pneuma:"0-10 — насколько фокус на дружбе, отношениях между друзьями"`
		RomanticTone    int    `pneuma:"0-10 — насколько выражена романтичность, чувства"`
		FlirtTone       int    `pneuma:"0-10 — насколько флиртовый характер"`
		AdventureLevel  int    `pneuma:"0-10 — дух приключений, путешествий, активности"`
		EmotionalDepth  int    `pneuma:"0-10 — насколько фраза затрагивает чувства, глубокие эмоции"`
		Theme           string `pneuma:"одно из: 'дружба', 'романтика', 'приключения', 'юмор', 'самооценка', 'ностальгия', 'школа', 'фантазии', 'другое'"`
	}

	d := ClassifiedPoll{}

	prompt := ` 
	- Ты — модератор тематических опросов в приложении для подростков.
	Проанализируй текст опроса и оцени его по заданным характеристикам, включая возможные признаки романтики, флирта, дружбы, эмоций и приключений.
	Учитывай подтекст, интонацию и возможную реакцию участников. Не оценивай допустимость — все опросы уместны.`

	poll := "Настоящий братан"

	p.FillWithContext(&d, poll+prompt)

	fmt.Println("Опрос: ", poll)
	fmt.Printf("%+v\n", d)

	d2 := ClassifiedPoll{}

	poll2 := "Такая милая сегодня"

	p.FillWithContext(&d2, poll2+prompt)

	fmt.Println("Опрос: ", poll2)
	fmt.Printf("%+v\n", d2)

}
