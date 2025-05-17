package models

import (
	"time"

	"gorm.io/gorm"
)

// Portfolio - Projects, skills, etc.
type Portfolio struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `json:"image_url"`
	ProjectURL  string `json:"project_url"`
	Skills      string `json:"skills" gorm:"type:text"`
}

// BlogPost - Blog articles
type BlogPost struct {
	gorm.Model
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Published time.Time `json:"published" gorm:"default:CURRENT_TIMESTAMP"`
	ImageURL  string    `json:"image_url"`
}

// Resume - Education, experience, etc.
type Resume struct {
	gorm.Model
	Category    string `json:"category" gorm:"not null"` // education, experience, skill
	Title       string `json:"title" gorm:"not null"`
	Institution string `json:"institution"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description" gorm:"type:text"`
}

// SocialMedia - Links to social media profiles
type SocialMedia struct {
	gorm.Model
	Platform string `json:"platform" gorm:"not null"`
	URL      string `json:"url" gorm:"not null"`
	Icon     string `json:"icon"`
}

// User - Admin user for authentication
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
