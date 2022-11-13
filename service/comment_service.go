package service

import (
	"errors"
	"project2/model/entity"
	"project2/model/input"
	"project2/repository"
)

type commentService struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

type CommentService interface {
	CreateComment(input input.CommentInput, idUser int) (entity.Comment, error)
	GetComment(UserID int) ([]entity.Comment, error)
	DeleteComment(ID int) (entity.Comment, error)
	UpdateComment(ID int, input input.CommentUpdateInput) (entity.Comment, error)
	GetCommentByID(commentID int) (entity.Comment, error)
	GetCommentsByPhotoID(photoID int) ([]entity.Comment, error)
}

func NewCommentService(commentRepository repository.CommentRepository, photoRepository repository.PhotoRepository) *commentService {
	return &commentService{commentRepository, photoRepository}
}

func (s *commentService) CreateComment(input input.CommentInput, idUser int) (entity.Comment, error) {
	photoData, err := s.photoRepository.FindByID(input.PhotoID)

	if err != nil {
		return entity.Comment{}, err
	}
	if photoData.ID == 0 {
		return entity.Comment{}, errors.New("photo not found")
	}

	newComment := entity.Comment{
		Message: input.Message,
		PhotoID: input.PhotoID,
		UserID:  idUser,
	}

	createNewcomment, err := s.commentRepository.Save(newComment)

	if err != nil {
		return entity.Comment{}, err
	}

	return createNewcomment, nil
}

func (s *commentService) GetComment(UserID int) ([]entity.Comment, error) {
	comment, err := s.commentRepository.FindByUserID(UserID)
	if err != nil {
		return []entity.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) DeleteComment(ID int) (entity.Comment, error) {
	comment, err := s.commentRepository.FindByID(ID)

	if err != nil {
		return entity.Comment{}, err
	}

	if comment.ID == 0 {
		return entity.Comment{}, nil
	}

	Deleted, err := s.commentRepository.Delete(ID)

	if err != nil {
		return entity.Comment{}, err
	}

	return Deleted, nil
}

func (s *commentService) UpdateComment(ID int, input input.CommentUpdateInput) (entity.Comment, error) {

	Result, err := s.commentRepository.FindByID(ID)

	if err != nil {
		return entity.Comment{}, err
	}

	if Result.ID == 0 {
		return entity.Comment{}, nil
	}

	updated := entity.Comment{
		Message: input.Message,
	}

	commentUpdate, err := s.commentRepository.Update(updated, ID)

	if err != nil {
		return entity.Comment{}, err
	}

	return commentUpdate, nil
}

func (s *commentService) GetCommentByID(commentID int) (entity.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) GetCommentsByPhotoID(photoID int) ([]entity.Comment, error) {
	comments, err := s.commentRepository.FindByPhotoID(photoID)

	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}
