package main

import (
	"tauth/app"
	"tauth/utils"
)

func main() {
	utils.Migrate()
	app.App()
}
