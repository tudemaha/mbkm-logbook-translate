package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/tudemaha/mbkm-logbook-translate/dto"
)

func GetLogbook(week int) (dto.Logbook, error) {
	req, err := http.NewRequest("GET", os.Getenv("BASE_URL")+fmt.Sprint(week), nil)
	if err != nil {
		return dto.Logbook{}, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("JWT"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return dto.Logbook{}, err
	}
	defer resp.Body.Close()

	body := json.NewDecoder(resp.Body)

	if resp.StatusCode != 200 {
		var errorBody dto.ErrorResponse
		err := body.Decode(&errorBody)
		if err != nil {
			return dto.Logbook{}, err
		}

		return dto.Logbook{}, errors.New(errorBody.Error.Message)
	}

	var logbook dto.Logbook
	err = body.Decode(&logbook)
	if err != nil {
		return dto.Logbook{}, err
	}

	return logbook, nil
}
