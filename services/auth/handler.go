package authsvc

import (
	"tauth/dbops/gorm/users"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

type AuthSvcImpl struct {
	userGorm users.GormInterface
}

// interface.
type Interface interface {
	EmailCheck(ctx *gin.Context, req EmailCheckRequest) (models.BaseResponse, bool, error)
	RegisterUser(ctx *gin.Context, req SingupObject) (models.BaseResponse, entities.Users, error)
	LoginUser(ctx *gin.Context, req LoginObject) (models.BaseResponse, entities.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                                LOGIN HANDLER                               */
/* -------------------------------------------------------------------------- */
func Handler(userGorm users.GormInterface) *AuthSvcImpl {
	return &AuthSvcImpl{
		userGorm: userGorm,
	}
}
