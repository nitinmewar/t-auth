package app

import (
	"tauth/app/middleware"
	"tauth/controllers/auth"
	"tauth/database"
	"tauth/dbops/gorm/users"
	authsvc "tauth/services/auth"

	"github.com/gin-gonic/gin"
)

// @todo discuss the nomenclature with jitin

func MapURL() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	gormDB, _ := database.Connection()

	usersGorm := users.Gorm(gormDB)
	authSvc := authsvc.Handler(usersGorm)
	authHanlder := auth.Handler(authSvc)

	router.POST("/email-check", authHanlder.EmailCheck)
	router.POST("/signup", authHanlder.Signup)
	router.POST("/login", authHanlder.Login)

	// verifications
	err := router.Run()
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
