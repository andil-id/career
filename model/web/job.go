package web

import (
	"mime/multipart"
	"time"
)

type Job struct {
	Id          string    `json:"id,omitempty"`
	CategoryId  string    `json:"category_id,omitempty"`
	CompanyLogo string    `json:"company_logo,omitempty"`
	CompanyName string    `json:"company_name,omitempty"`
	Location    string    `json:"location,omitempty"`
	Title       string    `json:"title,omitempty"`
	Type        string    `json:"type,omitempty"`
	Banner      []string  `json:"banner,omitempty"`
	Description string    `json:"description,omitempty"`
	Email       string    `json:"email,omitempty"`
	WebsiteUrl  string    `json:"website_url,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CreateJob struct {
	CategoryId  string                  `form:"category_id" binding:"required"`
	CompanyLogo *multipart.FileHeader   `form:"company_logo" binding:"required"`
	CompanyName string                  `form:"company_name" binding:"required"`
	Location    string                  `form:"location" binding:"required"`
	Title       string                  `form:"title" binding:"required"`
	Type        string                  `form:"type" binding:"required"`
	Banner      []*multipart.FileHeader `form:"banner" binding:"required"`
	Description string                  `form:"description" binding:"required"`
	Email       string                  `form:"email" binding:"required"`
	WebsiteUrl  string                  `form:"website_url" binding:"required"`
}
