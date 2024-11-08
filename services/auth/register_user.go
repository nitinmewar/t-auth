package authsvc

import (
	"net/http"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* ------------------------------ REGISTER USER ----------------------------- */
func (h *AuthSvcImpl) RegisterUser(ctx *gin.Context, req SingupObject) (models.BaseResponse, entities.Users, error) {
	var res models.BaseResponse
	var user entities.Users
	var err error

	// default response
	res.Message = "something went wrong"
	res.StatusCode = http.StatusInternalServerError

	// hash the password
	pass := []byte(req.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	// map the data
	user.PrimaryEmail = req.Email
	user.Password = string(hash)

	// create a user
	user, err = h.userGorm.CreateUser(ctx, user)
	if err != nil {
		return res, user, err
	}

	// success response
	res.Success = true
	res.Message = "signup succcesfull"
	res.StatusCode = http.StatusOK

	return res, user, err
}
