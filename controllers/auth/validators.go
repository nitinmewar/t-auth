package auth

import (
	"tauth/database"
	"tauth/dbops/gorm/users"
	authsvc "tauth/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

/* ------------------------- EMAIL CHECK VALIDATION ------------------------- */
func validateEmailCheckReq(ctx *gin.Context) (authsvc.EmailCheckRequest, error) {
	var reqBody authsvc.EmailCheckRequest
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}
	validate := validator.New()

	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	if reqBody.Email == "" {
		return reqBody, errors.New("email field is mandatory")
	}

	return reqBody, err
}

/* -------------------------------------------------------------------------- */
/*                                  REGISTER                                  */
/* -------------------------------------------------------------------------- */
func validateRegisterReq(ctx *gin.Context) (authsvc.SingupObject, error) {
	var reqBody authsvc.SingupObject
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}
	validate := validator.New()

	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	if reqBody.Email == "" || reqBody.Password == "" {
		return reqBody, errors.New("email and password is required")
	}

	db, _ := database.Connection()
	usersGorm := users.Gorm(db)

	exist, err := usersGorm.FindEmail(ctx, reqBody.Email)
	if err != nil {
		return reqBody, err
	}

	if exist {
		return reqBody, errors.New("email already registered")
	}

	return reqBody, err
}

/* ------------------------- VALIDATE LOGIN REQUEST ------------------------- */
func validateLoginReq(ctx *gin.Context) (authsvc.LoginObject, error) {
	var reqBody authsvc.LoginObject
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}
	validate := validator.New()

	err = validate.Struct(reqBody)
	if err != nil {
		return reqBody, err
	}

	if reqBody.Email == "" || reqBody.Password == "" {
		return reqBody, errors.New("email and password is required")
	}

	db, _ := database.Connection()
	usersGorm := users.Gorm(db)

	exist, err := usersGorm.FindEmail(ctx, reqBody.Email)
	if err != nil {
		return reqBody, err
	}

	if !exist {
		return reqBody, errors.New("email does not exist")
	}

	return reqBody, err
}
