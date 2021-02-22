package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func main() {
	cat := catalog.NewBuilder()
	cat.Set(language.English, "You are %d minute(s) late.",
		plural.Selectf(1, "",
			plural.One, "You are 1 minute late.",
			plural.Other, "You are %d minutes late.",
		))
	cat.Set(language.Russian, "You are %d minute(s) late.",
		plural.Selectf(1, "",
			plural.One, "Вы опоздали на 1 минуту.",
			plural.Few, "Вы опоздали на %d минуты.",
			plural.Other, "Вы опоздали на %d минут.",
		))

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d: ", i)
		p := message.NewPrinter(language.Russian, message.Catalog(cat))
		p.Printf("You are %d minute(s) late.", i)
		fmt.Println()
	}
}
