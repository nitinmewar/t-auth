package tautherrors

import (
	"net/http"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                       SERVICE UNAVAILABLE ERROR 503                        */
/* -------------------------------------------------------------------------- */
func ServiceUnavailable(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusServiceUnavailable
	smerror.Code = errorCode
	smerror.Type = errorType.ServiceUnavailable
	smerror.Message = "service unavailable, please try again after some time"
	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
