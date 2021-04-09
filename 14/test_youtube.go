package main

import (
	"github.com/kkdai/youtube"
	"log"
)

var (
	pathDir = "./videos"
)

func main() {
	url := "https://www.youtube.com/watch?v=nmDFmI2oNBY"
	yt := youtube.NewYoutube(false, true)
	err := yt.DecodeURL(url)
	if err != nil {
		log.Fatal(err)
	}
	fileName := yt.Author + "_" + yt.Title + ".mp4"
	err = yt.StartDownload(pathDir, fileName, "high", 0)
	if err != nil {
		log.Fatal(err)
	}
}
