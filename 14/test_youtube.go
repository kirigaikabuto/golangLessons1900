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
	err = yt.StartDownload(pathDir, "1.mp4", "high", 0)
	if err != nil {
		log.Fatal(err)
	}
}
