package main

import "time"

type User struct {
	ID        uint        `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty" sql:"DEFAULT:current_timestamp"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty" sql:"DEFAULT:current_timestamp on update current_timestamp"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
	Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
}
