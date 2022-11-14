package entity

import (
	"time"
)

type User struct {
	ID        int       
	Username  string    
	Email     string    
	Password  string    
	Age       int    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
