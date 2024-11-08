package tautherrors

import (
	"net/http"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            INTERNAL SERVER ERROR                           */
/* -------------------------------------------------------------------------- */
func InternalServer(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusInternalServerError
	smerror.Code = errorCode
	smerror.Type = errorType.server
	smerror.Message = "something went wrong"
	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
