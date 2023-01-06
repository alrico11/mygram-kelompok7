package entity

import "time"

type SocialMedia struct {
	ID        int  `gorm:"primaryKey"`
	Name      string  
	URL       string    
	UserID    int     
	CreatedAt time.Time 
	UpdatedAt time.Time 
	User      User `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (SocialMedia) TableName() string {
	return "socialmedia"
}
