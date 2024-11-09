package keystrokesvc

import (
	"tauth/dbops/gorm/keystrokes"
	"tauth/dbops/gorm/users"
	"tauth/entities"
	"tauth/models"

	"github.com/gin-gonic/gin"
)

type KeyStrokeImpl struct {
	keyStrokeGorm keystrokes.GormInterface
	userGorm      users.GormInterface
}

// interface.
type Interface interface {
	CreateKeyStroke(ctx *gin.Context, req RequestBody) (models.BaseResponse, entities.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                                LOGIN HANDLER                               */
/* -------------------------------------------------------------------------- */
func Handler(keyStrokeGorm keystrokes.GormInterface, userGorm users.GormInterface) *KeyStrokeImpl {
	return &KeyStrokeImpl{
		keyStrokeGorm: keyStrokeGorm,
		userGorm:      userGorm,
	}
}
