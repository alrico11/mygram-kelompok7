package repository

import (
	"project2/model/entity"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Save(socialmedia entity.SocialMedia) (entity.SocialMedia, error)
	FindByUserID(ID int) ([]entity.SocialMedia, error)
	Update(socialmedia entity.SocialMedia, ID int) (entity.SocialMedia, error)
	Delete(ID int) (entity.SocialMedia, error)
	FindByID(ID int) (entity.SocialMedia, error)
}

type socialmediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialmediaRepository {
	return &socialmediaRepository{db}
}

func (r *socialmediaRepository) Save(socialmedia entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Create(&socialmedia).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialmediaRepository) FindByID(ID int) (entity.SocialMedia, error) {
	socialmedia := entity.SocialMedia{}

	err := r.db.Where("id = ?", ID).Find(&socialmedia).Error

	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *socialmediaRepository) FindByUserID(ID int) ([]entity.SocialMedia, error) {
	var socialmedia []entity.SocialMedia

	err := r.db.Where("User_id = ?", ID).Find(&socialmedia).Error

	if err != nil {
		return []entity.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialmediaRepository) Update(socialmedia entity.SocialMedia, ID int) (entity.SocialMedia, error) {
	err := r.db.Where("id = ?", ID).Updates(&socialmedia).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialmediaRepository) Delete(ID int) (entity.SocialMedia, error) {
	socialDeleted := entity.SocialMedia{}

	err := r.db.Where("id = ?", ID).Delete(&socialDeleted).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialDeleted, err
}
