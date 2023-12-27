package main

import (
	"flag"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tudemaha/mbkm-logbook-translate/internal/translation/controller"
)

func main() {
	var sourceLanguage, targetLanguage, weeks string
	flag.StringVar(&sourceLanguage, "source", "en-US", "source language of your logbook")
	flag.StringVar(&targetLanguage, "target", "id-ID", "target language of your translated logbook")
	flag.StringVar(&weeks, "weeks", "1", "number of weeks to be translated (separated by comma)")

	flag.Parse()

	weekSlice := strings.Split(weeks, ",")
	for _, w := range weekSlice {
		weekInt, _ := strconv.Atoi(w)
		controller.GetTranslationPerWeek(sourceLanguage, targetLanguage, weekInt)
	}
}
