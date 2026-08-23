package main

import (
	"errors"
)

var ErrExistingUser = errors.New("error: user is existing")

type User struct {
	Email string
}

type UserStore interface {
	GetUser(email string) *User
	Save(email string) error
}

type Mailer interface {
	SendWelcome(email string) error
}

type RegistrationService struct {
	userStore UserStore
	mailer    Mailer
}

func NewRegistrationService(store UserStore, mailer Mailer) *RegistrationService {
	return &RegistrationService{
		userStore: store,
		mailer:    mailer,
	}
}

func (rs *RegistrationService) Register(email string) error {
	err := rs.userStore.Save(email)
	if err != nil {
		return err
	}

	err = rs.mailer.SendWelcome(email)
	if err != nil {
		return err
	}

	return nil
}
