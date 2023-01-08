// Package main runs the shopping list
package main

import (
	"flag"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gowhale/go-karaoke/pkg/config"
	gc "github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
	"github.com/gowhale/led-matrix-golang/pkg/matrix"
)

func main() {
	lyrics, err := config.LoadLyrics("isaiah-dreads-ratings-freestyle.json")
	if err != nil {
		log.Fatalln(err)
	}

	startTimes := []int{}
	for _, lyric := range lyrics.Lyrics {
		startTimes = append(startTimes, lyric.StartTime)
	}

	var debugMode = flag.Bool("debug", false, "run in debug mode")
	flag.Parse()

	cfg := gc.PinConfig{
		RowPins: []int{1, 2, 3, 4, 5},
		ColPins: make([]int, 180),
	}

	scrn := gui.NewTerminalGui(cfg)
	if !*debugMode {
		var err error
		scrn, err = gui.NewledGUI(cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := scrn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := startKaraoke(scrn, lyrics, cfg, startTimes); err != nil {
		log.Fatalln(err)
	}
}

func startKaraoke(scrn gui.Screen, lyrics config.LyricsConfig, cfg gc.PinConfig, startTimes []int) error {
	startTime := time.Now()
	if err := scrn.AllLEDSOff(); err != nil {
		return err
	}
	count := 0
	for count < len(startTimes)-1 {
		nextStartTime := (time.Duration(startTimes[count]) * time.Second).Seconds()
		currentTime := time.Duration(time.Since(startTime)).Seconds()
		if nextStartTime < currentTime {
			if err := displayLyric(scrn, lyrics, cfg, startTimes, count); err != nil {
				return err
			}

			count++
		}
	}
	return nil
}

func displayLyric(scrn gui.Screen, lyrics config.LyricsConfig, cfg gc.PinConfig, startTimes []int, count int) error {
	displayTime := startTimes[count+1] - startTimes[count]

	lowercaseString := strings.ToLower(lyrics.Lyrics[count].Line)
	lowercaseString = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(lowercaseString, "")
	matrixLyric, err := matrix.ConcatanateLetters(lowercaseString)
	if err != nil {
		return err
	}
	matrixLyric, err = matrix.TrimMatrix(matrixLyric, cfg, 0)
	if err != nil {
		return err
	}

	if err := scrn.DisplayMatrix(matrixLyric, time.Second*time.Duration(displayTime)); err != nil {
		return err
	}
	return scrn.AllLEDSOff()
}
