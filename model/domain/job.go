package domain

import "time"

type Job struct {
	Id          string
	CategoryId  string
	CompanyLogo string
	CompanyName string
	Location    string
	Title       string
	Type        string
	Banner      string
	Description string
	Email       string
	WebsiteUrl  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
