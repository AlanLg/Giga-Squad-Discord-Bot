package models

import (
	"fmt"
)

type Game struct {
	Appid           string `json:"appid"`
	Name            string `json:"name"`
	Playtime2Weeks  int    `json:"playtime_2weeks"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconUrl      string `json:"img_icon_url"`
}

type RecentlyPlayedGames struct {
	TotalCount int    `json:"totalCount"`
	Games      []Game `json:"games"`
}

type Response struct {
	Response RecentlyPlayedGames `json:"response"`
}

func (recentlyPlayedGames RecentlyPlayedGames) String() string {
	result := ""

	for _, game := range recentlyPlayedGames.Games {
		result += game.String()
	}

	return fmt.Sprintf("Steam Games Played Recently:\n%s", result)
}

func (game Game) String() string {
	return fmt.Sprintf("Game: %s\nHours Played: %dh\n\n", game.Name, game.PlaytimeForever/60)
}

func (response Response) String() string {
	return response.Response.String()
}
