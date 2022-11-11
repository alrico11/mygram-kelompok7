package repository

import (
	"project2/model/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Save(photo entity.Photo) (entity.Photo, error)
	Delete(ID int) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	FindByID(ID int) (entity.Photo, error)
	FindByUserID(ID int) ([]entity.Photo, error)
	Update(photo entity.Photo, ID int) (entity.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Save(photo entity.Photo) (entity.Photo, error) {
	err := r.db.Save(&photo).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return photo, nil
}

func (r *photoRepository) GetAll() ([]entity.Photo, error) {
	var photos []entity.Photo

	err := r.db.Find(&photos).Error

	if err != nil {
		return []entity.Photo{}, err
	}

	return photos, nil
}

func (r *photoRepository) FindByID(ID int) (entity.Photo, error) {
	photo := entity.Photo{}

	err := r.db.Where("id = ?", ID).Find(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) FindByUserID(ID int) ([]entity.Photo, error) {
	var photos []entity.Photo

	err := r.db.Where("user_id = ?", ID).Find(&photos).Error

	if err != nil {
		return []entity.Photo{}, err
	}

	return photos, nil
}

func (r *photoRepository) Delete(ID int) (entity.Photo, error) {
	photoDeleted := entity.Photo{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&photoDeleted).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return photoDeleted, err
}

func (r *photoRepository) Update(photo entity.Photo, ID int) (entity.Photo, error) {
	err := r.db.Where("id = ?", ID).Updates(&photo).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return photo, nil
}
