package main

import (
	"tauth/app"
	"tauth/config"
	"tauth/utils"
)

func main() {
	config.LoadConfigs()
	utils.Migrate()
	app.App()
}
