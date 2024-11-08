package tautherrors

import (
	"net/http"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

/* -------------------------------------------------------------------------- */
/*                            VALIDATION ERROR 422                            */
/* -------------------------------------------------------------------------- */
func Validation(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusUnprocessableEntity

	smerror.Code = errorCode
	smerror.Type = errorType.validation
	smerror.Message = err

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
