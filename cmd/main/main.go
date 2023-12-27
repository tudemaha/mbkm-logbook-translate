package main

import (
	"flag"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var sourceLanguage, targetLanguage, weeks string
	flag.StringVar(&sourceLanguage, "source", "en-US", "source language of your logbook")
	flag.StringVar(&targetLanguage, "target", "id-ID", "target language of your translated logbook")
	flag.StringVar(&weeks, "weeks", "1", "number of weeks to be translated (separated by comma)")

	flag.Parse()
}
