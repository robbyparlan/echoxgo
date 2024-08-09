package user

import (
	"time"
)

type UserEntity struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
	Email     string
	Fullname  string
	IsActive  string
	Roles     int
}
