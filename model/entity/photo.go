package entity

import (
	"time"
)

type Photo struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Caption   string 
	PhotoURL  string
	UserID    int   
	CreatedAt time.Time 
	UpdatedAt time.Time 
	User      User `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Photo) TableName() string {
	return "photo"
}
