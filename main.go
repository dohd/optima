package main

import (
	"optima-app/config"
	"optima-app/internal/app"
)

func main() {
	config.InitDatabase()
	app.Run()
}
