package main

import (
	"fmt"
	"github.com/kkdai/youtube"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	pathDir = ""
	url     = ""
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

func mainAction(c *cli.Context) error {
	fmt.Println("Welcome")
	return nil
}

func downloadAction(c *cli.Context) error {
	err := downloadFromYoutube(url)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path,p",
			Value:       "./videos",
			Destination: &pathDir,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "for download video",
			Action:  downloadAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "url,u",
					Usage:       "url string",
					Destination: &url,
					Value:       "https://www.youtube.com/watch?v=nmDFmI2oNBY",
				},
			},
		},
	}
	app.Action = mainAction
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
