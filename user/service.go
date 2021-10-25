package user

import "errors"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	SignIn(input SignInInput) (User, error)
	IsEmailExist(email string) (bool, error)
	GetUserByUuid(uuid string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{
		Uuid:          input.Uuid,
		Email:         input.Email,
		VerifiedEmail: input.VerifiedEmail,
		Picture:       input.Picture,
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) SignIn(input SignInInput) (User, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Uuid == "" {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (s *service) IsEmailExist(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Uuid != "" {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByUuid(uuid string) (User, error) {
	user, err := s.repository.FindByUuid(uuid)
	if err != nil {
		return user, err
	}

	if user.Uuid == "" {
		return user, errors.New("no user found with that id")
	}

	return user, nil
}
