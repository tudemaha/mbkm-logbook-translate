package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateJson(week string, jsonData interface{}) error {
	jsonByte, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("./translated/%s.json", week), jsonByte, 0644)
	if err != nil {
		return err
	}

	return nil
}
