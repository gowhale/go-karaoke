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

	startTime := time.Now()

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

	scrn.AllLEDSOff()
	count := 0
	for count < len(startTimes)-1 {
		nextStartTime := (time.Duration(startTimes[count]) * time.Second).Seconds()
		currentTime := time.Duration(time.Since(startTime)).Seconds()
		if nextStartTime < currentTime {
			displayTime := startTimes[count+1] - startTimes[count]
			// log.Println(currentTime, displayTime, lyrics.Lyrics[count])

			lowercaseString := strings.ToLower(lyrics.Lyrics[count].Line)
			lowercaseString = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(lowercaseString, "")
			matrixLyric, err := matrix.ConcatanateLetters(lowercaseString)
			if err != nil {
				log.Fatalln(err)
			}
			matrixLyric, err = matrix.TrimMatrix(matrixLyric, cfg, 0)
			if err != nil {
				log.Fatalln(err)
			}

			scrn.DisplayMatrix(matrixLyric, time.Second*time.Duration(displayTime))
			scrn.AllLEDSOff()
			count++
		}
	}
}
