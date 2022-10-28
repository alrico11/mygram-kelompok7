package service

import (
	"project2/model/entity"
	"project2/model/input"
	"project2/repository"
)

type SocialMediaService interface {
	CreateSocialMedia(input input.SocialInput, idUser int) (entity.SocialMedia, error)
	DeleteSocialMedia(ID int) (entity.SocialMedia, error)
	UpdateSocialMedia(ID int, input input.SocialInput) (entity.SocialMedia, error)
	GetSocialMedia(UserID int) ([]entity.SocialMedia, error)
}
type socialmediaService struct {
	socialmediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialmediaRepository repository.SocialMediaRepository) *socialmediaService {
	return &socialmediaService{socialmediaRepository}
}

func (s *socialmediaService) CreateSocialMedia(input input.SocialInput, idUser int) (entity.SocialMedia, error) {
	newSocialMedia := entity.SocialMedia{
		Name:   input.Name,
		URL:    input.URL,
		UserID: idUser,
	}

	createdSocialmedia, err := s.socialmediaRepository.Save(newSocialMedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return createdSocialmedia, nil

}

func (s *socialmediaService) GetSocialMedia(UserID int) ([]entity.SocialMedia, error) {
	socialmedia, err := s.socialmediaRepository.FindByUserID(UserID)

	if err != nil {
		return []entity.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *socialmediaService) DeleteSocialMedia(ID int) (entity.SocialMedia, error) {
	socialmedia, err := s.socialmediaRepository.FindByID(ID)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	if socialmedia.ID == 0 {
		return entity.SocialMedia{}, nil
	}

	socialmediaDeleted, err := s.socialmediaRepository.Delete(ID)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmediaDeleted, nil
}

func (s *socialmediaService) UpdateSocialMedia(ID int, input input.SocialInput) (entity.SocialMedia, error) {

	Result, err := s.socialmediaRepository.FindByID(ID)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	if Result.ID == 0 {
		return entity.SocialMedia{}, nil
	}

	updated := entity.SocialMedia{
		Name: input.Name,
		URL:  input.URL,
	}

	socialmediaUpdate, err := s.socialmediaRepository.Update(updated, ID)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmediaUpdate, nil
}
