package repository

import (
	"project2/model/entity"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

type CommentRepository interface {
	Save(comment entity.Comment) (entity.Comment, error)
	Delete(ID int) (entity.Comment, error)
	FindByUserID(ID int) ([]entity.Comment, error)
	Update(comment entity.Comment, ID int) (entity.Comment, error)
	FindByID(ID int) (entity.Comment, error)
	FindByPhotoID(IDPhoto int) ([]entity.Comment, error)
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Save(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Preload("User").Preload("Photo").Save(&comment).Error

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) Delete(ID int) (entity.Comment, error) {
	commentDeleted := entity.Comment{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&commentDeleted).Error

	if err != nil {
		return entity.Comment{}, err
	}

	return commentDeleted, err
}

func (r *commentRepository) FindByUserID(ID int) ([]entity.Comment, error) {
	var comment []entity.Comment

	err := r.db.Preload("User").Preload("Photo").Where("user_id = ?", ID).Find(&comment).Error

	if err != nil {
		return []entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) FindByID(ID int) (entity.Comment, error) {
	comment := entity.Comment{}

	err := r.db.Where("id = ?", ID).Find(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) Update(comment entity.Comment, ID int) (entity.Comment, error) {
	err := r.db.Where("id = ?", ID).Updates(&comment).Error

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) FindByPhotoID(IDPhoto int) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Where("photo_id = ?", IDPhoto).Find(&comments).Error

	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}
