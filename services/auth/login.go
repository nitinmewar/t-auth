package authsvc

import (
	"net/http"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* ---------------------------------- LOGIN --------------------------------- */
func (h *AuthSvcImpl) LoginUser(ctx *gin.Context, req LoginObject) (models.BaseResponse, entities.Users, error) {
	var res models.BaseResponse
	var user entities.Users
	var err error

	// default response
	res.Message = "something went wrong"
	res.StatusCode = http.StatusInternalServerError

	// get user
	user, err = h.userGorm.GetUserDetailsEmail(ctx, req.Email)
	if err != nil {
		return res, user, err
	}

	// check password
	err = h.checkPassword(ctx, req)
	if err != nil {
		res.Success = false
		res.Message = "incorrect password"
		res.StatusCode = http.StatusUnauthorized
		return res, user, nil
	}

	// success response
	res.Success = true
	res.Message = "login succcesfull"
	res.StatusCode = http.StatusOK

	return res, user, err
}

func (h *AuthSvcImpl) checkPassword(c *gin.Context, req LoginObject) error {
	res, err := h.userGorm.GetUserDetailsEmail(c, req.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return err
	}

	return nil
}
