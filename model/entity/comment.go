package entity

import (
	"time"
)

type Comment struct {
	ID        int  `gorm:"primaryKey"`
	UserID    int      
	PhotoID   int  
	Message   string  
	CreatedAt time.Time 
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photo     Photo `gorm:"foreignKey:PhotoID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Comment) TableName() string {
	return "comment"
}
