package handlers

import (
	"encoding/json"
	"fmt"
	"giga-squad-bot/models"
	"io"
	"log"
	"net/http"
)

func GetRecentlyPlayedGames() string {
	resp, err := http.Get("")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var response models.Response
	json.Unmarshal(bodyBytes, &response)

	fmt.Println(response.String())
	return fmt.Sprintf("%s", response)
}
