package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Lyric struct {
	Line      string `json:"line"`
	StartTime int    `json:"start"`
}

type LyricsConfig struct {
	Lyrics []Lyric `json:"lyrics"`
}

func LoadLyrics(filename string) (LyricsConfig, error) {
	filepath := fmt.Sprintf("lyrics/%s", filename)

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return LyricsConfig{}, err
	}

	var lyricsConfig LyricsConfig

	json.Unmarshal(byteValue, &lyricsConfig)
	return lyricsConfig, nil
}
