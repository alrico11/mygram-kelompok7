package entity

import "time"

type SocialMedia struct {
	ID        int    
	Name      string  
	URL       string    
	UserID    int     
	CreatedAt time.Time 
	UpdatedAt time.Time 
	User      User
}

func (SocialMedia) TableName() string {
	return "socialmedia"
}
