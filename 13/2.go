package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)
//go mod init
//go get github.com/urfave/cli
//go run 2.go -> hello from cli app
//go run 2.go help
func main() {
	//создание приложения
	app := cli.NewApp()
	//указываем свойства
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello from cli app")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}
