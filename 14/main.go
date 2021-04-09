package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	path string
	a    int
	b    int
)

func mainAction(c *cli.Context) error {
	fmt.Println("Hello from main action", path)
	return nil
}

func downloadAction(c *cli.Context) error {
	fmt.Println("Hello from download action", path)
	return nil
}

func plusAction(c *cli.Context) error {
	fmt.Println(a + b)
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
			Value:       "default path",
			Destination: &path,
		},
		cli.IntFlag{
			Name:        "a",
			Value:       0,
			Destination: &a,
		},
		cli.IntFlag{
			Name:        "b",
			Value:       0,
			Destination: &b,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "download video by link",
			Action:  downloadAction,
		},
		{
			Name:    "plus",
			Aliases: []string{"+"},
			Usage:   "for calculate two numbers",
			Action:  plusAction,
		},
	}
	app.Action = mainAction
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
