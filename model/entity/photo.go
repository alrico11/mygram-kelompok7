package entity

import (
	"time"
)

type Photo struct {
	ID        int
	Title     string
	Caption   string 
	PhotoURL  string
	UserID    int   
	CreatedAt time.Time 
	UpdatedAt time.Time 
	User      User
}

func (Photo) TableName() string {
	return "photo"
}
