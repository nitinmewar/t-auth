package users

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
	CreateUser(ctx *gin.Context, user entities.Users) (entities.Users, error)
	GetUserDetailsEmail(ctx *gin.Context, email string) (entities.Users, error)
	FindEmail(ctx *gin.Context, email string) (res bool, err error)
	GetUserDetailsByPID(ctx *gin.Context, pid string) (entities.Users, error)
	UpdateKeystrokeMetrics(ctx *gin.Context, pid string) (entities.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *usersGormImpl {
	return &usersGormImpl{
		DB: gormDB,
	}
}

type usersGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */
func (r *usersGormImpl) CreateUser(ctx *gin.Context, user entities.Users) (entities.Users, error) {

	user.PID = sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.Users), Valid: true}
	err := r.DB.Session(&gorm.Session{}).Create(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "[CreateUser][Create]")
	}
	return user, nil
}

func (r *usersGormImpl) FindEmail(ctx *gin.Context, email string) (bool, error) {
	res := false
	var user entities.Users
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("primary_email = ?", email).
		Scopes(dbops.DeletedScopes(ctx)).
		Find(&user)
	err := result.Error
	if err != nil {
		return res, err
	}
	if result.RowsAffected != 0 {
		res = true
	}
	return res, nil
}

func (r *usersGormImpl) GetUserDetailsByPID(ctx *gin.Context, pid string) (entities.Users, error) {
	var user entities.Users
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&user)
	err := result.Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersGormImpl) GetUserDetailsEmail(ctx *gin.Context, email string) (entities.Users, error) {
	var user entities.Users
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("primary_email = ?", email).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&user)
	err := result.Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersGormImpl) UpdateKeystrokeMetrics(ctx *gin.Context, pid string) (entities.Users, error) {
	var user entities.Users
	db := r.DB.Session(&gorm.Session{})
	result := db.Model(&user).Where("user_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		Update("is_keystroke_calculated", true).
		Take(&user)
	err := result.Error
	if err != nil {
		return user, err
	}

	return user, nil
}
