package main

import (
	"video-tutorials/app"
)

func main() {
	a := app.New()
	a.InitDB()
	a.Start()
}
