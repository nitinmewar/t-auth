package keystroke

import (
	"errors"
	"tauth/services/keystrokesvc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func validateKeystrokeReq(ctx *gin.Context) (keystrokesvc.RequestBody, error) {
	var reqBody keystrokesvc.RequestBody
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

	if reqBody.Data.UserPID == "" {
		return reqBody, errors.New("user id is mandatory")
	}

	return reqBody, err
}
