package users

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"tauth/config"
	"tauth/constants"
	"tauth/database"
	"tauth/entities"
	"tauth/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                                 Create User                                */
/* -------------------------------------------------------------------------- */
func TestCreateUser(t *testing.T) {
	config.LoadConfigs()
	db, _ := database.Connection()
	usersGorm := Gorm(db)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	var user entities.Users

	user.PID = sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.Users), Valid: true}
	user.FirstName = "akhil"
	user.LastName = "babu"
	user.PrimaryEmail = "abc@xyz.com"
	user.Password = "nhujio"
	user.IsDeleted = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	res, err := usersGorm.CreateUser(c, user)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.PID)
	assert.NotEmpty(t, res.CreatedAt)
	assert.NotEmpty(t, res.UpdatedAt)

	assert.Equal(t, user.FirstName, res.FirstName)
	assert.Equal(t, user.IsDeleted, res.IsDeleted)
	assert.Equal(t, user.LastName, res.LastName)

	// test record cleanup
	t.Cleanup(func() {
		db.Model(&entities.Users{}).Where("user_pid = ?", res.PID).Delete(&res)
	})
}

/* -------------------------------------------------------------------------- */
/*                                 Find Email                                 */
/* -------------------------------------------------------------------------- */
func TestFindEmail(t *testing.T) {
	config.LoadConfigs()
	db, _ := database.Connection()
	usersGorm := Gorm(db)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}
	email := "jitin.bahri@gmail.com"

	res, err := usersGorm.FindEmail(c, email)
	assert.Empty(t, err)
	assert.Equal(t, true, res)
}
