// Package main runs the shopping list
package main

import (
	"log"
	"time"

	"github.com/gowhale/go-karaoke/pkg/config"
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

	// var debugMode = flag.Bool("debug", false, "run in debug mode")
	// flag.Parse()

	// cfg := "thirteen-by-five.json"
	// scrn := gui.NewTerminalGui(cfg)
	// if !*debugMode {
	// 	var err error
	// 	scrn, err = gui.NewledGUI(cfg)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// defer func() {
	// 	err := scrn.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	count := 0
	for count < len(startTimes)-1 {
		nextStartTime := (time.Duration(startTimes[count]) * time.Second).Seconds()
		currentTime := time.Duration(time.Since(startTime)).Seconds()
		if nextStartTime < currentTime {
			displayTime := startTimes[count+1] - startTimes[count]
			log.Println(currentTime, displayTime, lyrics.Lyrics[count])
			count++
		}
	}
}
