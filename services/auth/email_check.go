package authsvc

import (
	"fmt"
	"net/http"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

/* ------------------------------- EMAIL CHECK ------------------------------ */
func (h *AuthSvcImpl) EmailCheck(ctx *gin.Context, req EmailCheckRequest) (models.BaseResponse, bool, error) {
	var res models.BaseResponse
	var exist bool
	var err error

	// default response
	res.Message = "something went wrong"
	res.StatusCode = http.StatusInternalServerError

	exist, err = h.userGorm.FindEmail(ctx, req.Email)
	fmt.Println(err)
	if err != nil {
		return res, exist, err
	}

	// success response
	res.Success = true
	res.Message = "email check succcesfull"
	res.StatusCode = http.StatusOK

	return res, exist, err
}
