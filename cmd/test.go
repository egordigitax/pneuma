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

	type Dog struct {
		Name          string `pneuma:"use russian language"`
		Age           int    `pneuma:"more than 6"`
		FavouriteFood string `pneuma:"more like fruits"`
	}

	d := Dog{}

	p.Fill(&d)

	fmt.Println(d)
}
