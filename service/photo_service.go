package service

import (
	"errors"
	"project2/model/entity"
	"project2/model/input"
	"project2/repository"
)

type PhotoService interface {
	CreatePhoto(photoInput input.PhotoCreateInput, idUser int) (entity.Photo, error)
	DeletePhoto(idPhoto int, idUser int) (entity.Photo, error)
	GetPhotosAll() ([]entity.Photo, error)
	GetPhotosUser(idUser int) ([]entity.Photo, error)
	GetPhotoByID(idPhoto int) (entity.Photo, error)
	UpdatePhoto(idUser int, idPhoto int, input input.PhotoUpdateInput) (entity.Photo, error)
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) *photoService {
	return &photoService{photoRepository}
}

func (s *photoService) CreatePhoto(input input.PhotoCreateInput, idUser int) (entity.Photo, error) {
	newPhoto := entity.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   idUser,
	}

	createNewPhoto, err := s.photoRepository.Save(newPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	return createNewPhoto, nil

}

func (s *photoService) GetPhotosUser(idUser int) ([]entity.Photo, error) {
	photos, err := s.photoRepository.FindByUserID(idUser)

	if err != nil {
		return []entity.Photo{}, err
	}

	return photos, nil
}

func (s *photoService) GetPhotosAll() ([]entity.Photo, error) {
	photos, err := s.photoRepository.GetAll()

	if err != nil {
		return []entity.Photo{}, err
	}

	return photos, nil
}

func (s *photoService) DeletePhoto(idPhoto int, idUser int) (entity.Photo, error) {
	photoQuery, err := s.photoRepository.FindByID(idPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	if photoQuery.ID == 0 {
		return entity.Photo{}, errors.New("photo not found")
	}

	if idUser != photoQuery.UserID {
		return entity.Photo{}, errors.New("can't delete other people's photo")
	}

	photoDeleted, err := s.photoRepository.Delete(idPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	return photoDeleted, nil
}

func (s *photoService) GetPhotoByID(idPhoto int) (entity.Photo, error) {
	photoQuery, err := s.photoRepository.FindByID(idPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	if photoQuery.ID == 0 {
		return entity.Photo{}, nil
	}

	return photoQuery, nil
}

func (s *photoService) UpdatePhoto(idUser int, idPhoto int, input input.PhotoUpdateInput) (entity.Photo, error) {

	photoResult, err := s.photoRepository.FindByID(idPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	if photoResult.ID == 0 {
		return entity.Photo{}, errors.New("photo not found")
	}

	if idUser != photoResult.UserID {
		return entity.Photo{}, errors.New("can't update other people's photos")
	}

	updatedPhoto := entity.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   photoResult.UserID,
	}

	photoUpdate, err := s.photoRepository.Update(updatedPhoto, idPhoto)

	if err != nil {
		return entity.Photo{}, err
	}

	return photoUpdate, nil
}
