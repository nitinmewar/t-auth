package dbops

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                Deleted Scopes                              */
/* -------------------------------------------------------------------------- */
func DeletedScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_deleted = ?", false)
	}
}
