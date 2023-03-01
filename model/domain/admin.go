package domain

import "time"

type Admin struct {
	Id        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
