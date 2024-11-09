package app

import (
	"tauth/app/middleware"
	"tauth/controllers/auth"
	"tauth/controllers/keystroke"
	"tauth/database"
	"tauth/dbops/gorm/keystrokes"
	"tauth/dbops/gorm/users"
	authsvc "tauth/services/auth"
	"tauth/services/keystrokesvc"

	"github.com/gin-gonic/gin"
)

func MapURL() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	gormDB, _ := database.Connection()

	// user
	usersGorm := users.Gorm(gormDB)
	authSvc := authsvc.Handler(usersGorm)
	authHanlder := auth.Handler(authSvc)

	// keystroke profile
	keystrokeGorm := keystrokes.Gorm(gormDB)
	keystrokeSvc := keystrokesvc.Handler(keystrokeGorm, usersGorm)
	keyStrokeHanlder := keystroke.Handler(keystrokeSvc)

	router.POST("/email-check", authHanlder.EmailCheck)
	router.POST("/signup", authHanlder.Signup)
	router.POST("/login", authHanlder.Login)

	router.POST("/metrics", keyStrokeHanlder.CreateKeyStroke)

	// verifications
	err := router.Run()
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
