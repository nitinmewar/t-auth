package tautherrors

import (
	"net/http"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            INTERNAL SERVER ERROR                           */
/* -------------------------------------------------------------------------- */
func Unauthorized(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusUnauthorized
	smerror.Code = errorCode
	smerror.Type = errorType.Unauthorized
	smerror.Message = err
	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
