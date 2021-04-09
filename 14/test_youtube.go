package main

import (
	"github.com/kkdai/youtube"
	"log"
)

var (
	pathDir = "./videos"
)

func downloadFromYoutube(url string) error {
	yt := youtube.NewYoutube(false, true)
	err := yt.DecodeURL(url)
	if err != nil {
		return err
	}
	fileName := yt.Author + "_" + yt.Title + ".mp4"
	err = yt.StartDownload(pathDir, fileName, "high", 0)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := "https://www.youtube.com/watch?v=nmDFmI2oNBY"
	err := downloadFromYoutube(url)
	if err != nil {
		log.Fatal(err)
	}
}
