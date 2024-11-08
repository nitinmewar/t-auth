package tautherrors

import (
	"tauth/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                              DOWNSTREAM ERROR                              */
/* -------------------------------------------------------------------------- */
func Downstream(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := 550
	smerror.Code = errorCode
	smerror.Type = errorType.Downstream
	smerror.Message = "provider is down. Please try again after some time"
	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
