package keystrokes

import (
	"database/sql"
	"tauth/constants"
	"tauth/dbops"
	"tauth/entities"
	"tauth/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type GormInterface interface {
	CreateKeyStroke(ctx *gin.Context, keystroke entities.KeystrokeProfile) (entities.KeystrokeProfile, error)
	GetKeyStrokeByUserPID(ctx *gin.Context, pid string) (entities.KeystrokeProfile, error)
	DeletKeyStroke(ctx *gin.Context, pid string) (entities.KeystrokeProfile, error)
	ListKeyStrokes(ctx *gin.Context) ([]entities.KeystrokeProfile, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *keyStrokeImpl {
	return &keyStrokeImpl{
		DB: gormDB,
	}
}

type keyStrokeImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */
func (r *keyStrokeImpl) CreateKeyStroke(ctx *gin.Context, keystroke entities.KeystrokeProfile) (entities.KeystrokeProfile, error) {

	keystroke.PID = sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.Users), Valid: true}
	err := r.DB.Session(&gorm.Session{}).Create(&keystroke).Error
	if err != nil {
		return keystroke, errors.Wrap(err, "[CreateKeyStroke][Create]")
	}
	return keystroke, nil
}

func (r *keyStrokeImpl) GetKeyStrokeByUserPID(ctx *gin.Context, pid string) (entities.KeystrokeProfile, error) {
	var keystroke entities.KeystrokeProfile

	db := r.DB.Session(&gorm.Session{})

	result := db.Model(&keystroke).Where("user_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&keystroke)
	err := result.Error
	if err != nil {
		return keystroke, err
	}

	return keystroke, nil
}

func (r *keyStrokeImpl) DeletKeyStroke(ctx *gin.Context, pid string) (entities.KeystrokeProfile, error) {
	var keystroke entities.KeystrokeProfile
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		Update("is_deleted", true)
	err := result.Error
	if err != nil {
		return keystroke, err
	}

	return keystroke, nil
}

func (r *keyStrokeImpl) ListKeyStrokes(ctx *gin.Context) ([]entities.KeystrokeProfile, error) {
	var keystrokes []entities.KeystrokeProfile

	db := r.DB.Session(&gorm.Session{})

	result := db.Model(&entities.KeystrokeProfile{}).
		Scopes(dbops.DeletedScopes(ctx)).
		Find(&keystrokes)
	err := result.Error
	if err != nil {
		return keystrokes, err
	}

	return keystrokes, nil
}
