package entities

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type KeystrokeProfile struct {
	ID                int             `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID               sql.NullString  `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	UserPID           string          `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	SampleText        string          `gorm:"column:sample_text;type:varchar(100);not null"`
	TextLength        int             `gorm:"column:text_length;not null"`
	DwellTimes        pq.Float64Array `gorm:"column:dwell_times;type:float[]"`
	FlightTimes       pq.Float64Array `gorm:"column:flight_times;type:float[]"`
	AverageDwellTime  float64         `gorm:"column:average dwell_time;type:decimal(10,3)"`
	AverageFlightTime float64         `gorm:"column:average_flight_time;type:decimal(10,3)"`
	TotalTime         float64         `gorm:"column:total_time;type:decimal(10,3)"`
	WordsPerMinute    float64         `gorm:"column:words_per_minute;type:decimal(10,3)"`
	DeviceInfo        string          `gorm:"column:device_info;type:text"`
	CreatedFrom       string          `gorm:"column:created_from;type:varchar(255)"`
	SuccessfulMatches int             `gorm:"column:successful_matches"`
	FailedMatches     int             `gorm:"column:failed_matches"`
	IsDeleted         bool            `gorm:"column:is_deleted;not null;default:false"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastUsed          time.Time
}
