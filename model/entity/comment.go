package entity

import (
	"time"
)

type Comment struct {
	ID        int      
	UserID    int      
	PhotoID   int  
	Message   string  
	CreatedAt time.Time 
	UpdatedAt time.Time
	User      User
	Photo     Photo
}

func (Comment) TableName() string {
	return "comment"
}
