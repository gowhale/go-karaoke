// Package config deals with the loading of lyric files
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Lyric contains text and the time it should be displayed
type Lyric struct {
	Line      string `json:"line"`
	StartTime int    `json:"start"`
}

// LyricsConfig contains multiple lyrics in order
type LyricsConfig struct {
	Lyrics []Lyric `json:"lyrics"`
}

// LoadLyrics will load a lyrics file and return a struct for it
func LoadLyrics(filename string) (LyricsConfig, error) {
	return loadLyricsImpl(filename, &readJSON{})
}

func loadLyricsImpl(filename string, jReader readerJSON) (LyricsConfig, error) {
	filepath := fmt.Sprintf("lyrics/%s", filename)

	jsonFile, err := jReader.Open(filepath)
	if err != nil {
		return LyricsConfig{}, err
	}

	byteValue, err := jReader.ReadAll(jsonFile)
	if err != nil {
		return LyricsConfig{}, err
	}

	var lyricsConfig LyricsConfig

	if err := jReader.Unmarshal(byteValue, &lyricsConfig); err != nil {
		return LyricsConfig{}, err
	}

	return lyricsConfig, nil
}

type readJSON struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name readerJSON -inpkg --filename read_json_mock.go
type readerJSON interface {
	ReadAll(r io.Reader) ([]byte, error)
	Open(name string) (*os.File, error)
	Unmarshal(data []byte, v interface{}) error
}

func (*readJSON) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (*readJSON) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (*readJSON) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
