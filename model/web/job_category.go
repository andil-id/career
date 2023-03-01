package web

import "time"

type JobCategory struct {
	Id        string
	Name      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
