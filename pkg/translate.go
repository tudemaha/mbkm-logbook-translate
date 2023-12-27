package pkg

import (
	"context"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
	"google.golang.org/api/option"
)

func GetTranslatedReport(sourceLang, targetLang string, reports []string) ([]string, error) {
	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx,
		option.WithCredentialsFile(os.Getenv("KEY_PATH")))
	if err != nil {
		return nil, err
	}
	defer c.Close()

	req := &translatepb.TranslateTextRequest{
		Contents:           reports,
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		Parent:             "projects/" + os.Getenv("PROJECT_ID"),
	}

	resp, err := c.TranslateText(ctx, req)
	if err != nil {
		return nil, err
	}
	result := resp.GetTranslations()

	var translatedReport []string
	for _, res := range result {
		translatedReport = append(translatedReport, res.TranslatedText)
	}

	return translatedReport, nil
}
