package main

import (
	"fmt"
	"os"

	"github.com/egordigitax/pneuma"
	"github.com/egordigitax/pneuma/internal/llm/providers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	provider := providers.NewOpenAIProvider(os.Getenv("OPENAI_KEY"))
	provider.SetModel("gpt-4o")
	p := pneuma.InitWithProvider(provider)

	type ClassifiedPoll struct {
		FriendshipLevel int    `pneuma:"0-10 — насколько фокус на дружбе, отношениях между друзьями"`
		RomanticTone    bool   `pneuma:"есть ли тут романтический подтекст"`
		FlirtTone       bool   `pneuma:"есть ли тут флирт"`
		AdventureLevel  int    `pneuma:"0-10 — дух приключений, путешествий, активности"`
		EmotionalDepth  int    `pneuma:"0-10 — насколько фраза затрагивает чувства, глубокие эмоции"`
		Theme           string `pneuma:"дружба or романтика or приключения or юмор or самооценка or ностальгия or школа or фантазии or другое"`
	}

	d := ClassifiedPoll{}

	prompt := ` 
	- Ты — модератор тематических опросов в приложении для подростков. 
	Проанализируй текст опроса и оцени его по заданным характеристикам, включая возможные признаки романтики, флирта, дружбы, эмоций и приключений.
	Учитывай подтекст, интонацию и возможную реакцию участников. Не оценивай допустимость — все опросы уместны.
	например, опрос: "Такие красивые глаза" - это точно комплимент девочке, поэтому там френдшип - 4, романтик - 9, флирт - 10, адвенчур - 0, 
	эмоциональный уровень - 7, тема - самооценка. Или опрос: "Всегда на тебя положусь" - это скорее всего про братанов опрос для мужчин - френдшип 10, романтик - 0, флирт - 0
	и тд

	тема самооценка - когда опрос поднимает самооценку человеку, например - красивые глаза / стала бы известной актрисой
	тема романтика - когда есть флирт и заигрывание 
	тема юмор - когда что-то забавное или в шутливой форме
	`

	poll := "Настоящий братан"

	p.FillWithContext(&d, poll+prompt)

	fmt.Println("Опрос: ", poll)
	fmt.Printf("%+v\n", d)

	d2 := ClassifiedPoll{}

	poll2 := "Стала бы известной певицей"

	p.FillWithContext(&d2, poll2+prompt)

	fmt.Println("Опрос: ", poll2)
	fmt.Printf("%+v\n", d2)

}
