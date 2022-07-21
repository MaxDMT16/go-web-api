package main

import "getheadway/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":9912")
}
