package user

import "time"

type User struct {
	Id            int `gorm:"primaryKey"`
	Uuid          string
	Email         string
	VerifiedEmail bool
	Picture       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
