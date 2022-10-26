package repository

import (
	"project2/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Update(ID int, user entity.User) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
	Delete(ID int) (entity.User, error)
	// Get(ID int) (entity.User, error)
	// GetAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	user := entity.User{}

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (entity.User, error) {
	user := entity.User{}

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(ID int, user entity.User) (entity.User, error) {
	err := r.db.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ID int) (entity.User, error) {
	userDeleted := entity.User{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&userDeleted).Error

	if err != nil {
		return entity.User{}, err
	}

	return userDeleted, err
}
