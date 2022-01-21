package main

import (
	"video-tutorials/app"
)

func main() {
	a := app.New()
	a.InitDBs()
	a.Start()
}
