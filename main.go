package main

import "github.com/MaxDMT16/go-web-api/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":9912")
}
