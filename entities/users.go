package entities

import (
	"database/sql"
	"time"
)

type Users struct {
	ID                    int            `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID                   sql.NullString `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	FirstName             string         `gorm:"column:first_name;type:varchar(40)"`
	LastName              string         `gorm:"column:last_name;type:varchar(40)"`
	PrimaryEmail          string         `gorm:"column:primary_email;not null;type:varchar(100)"`
	Password              string         `gorm:"column:password;not null;type:varchar(100)"`
	IsKeystrokeCalculated bool           `gorm:"column:is_keystroke_calculated;not null;default:false"`
	IsDeleted             bool           `gorm:"column:is_deleted;not null;default:false"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
