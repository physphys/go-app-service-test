package model

import (
	"errors"
	"github.com/google/uuid"
	"unicode/utf8"
)

type (
	UserName string
	UserID   string

	User struct {
		id   UserID
		name UserName
	}
)

func (u User) Name() UserName {
	return u.name
}

func (u User) ID() UserID {
	return u.id
}

func NewUserID() (UserID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return UserID(newUUID.String()), nil
}

func NewUserName(name string) (UserName, error) {
	if name == "" {
		return "", errors.New("empty user name")
	}

	nameLength := utf8.RuneCountInString(name)
	const minNameLength = 3
	const maxNameLength = 20
	if nameLength < minNameLength {
		return "", errors.New("user name must be over than 3 chars ")
	}

	if nameLength > maxNameLength {
		return "", errors.New("user name must be less than 20 chars ")
	}

	return UserName(name), nil
}

func NewUser(name string) (User, error) {
	newUserID, err := NewUserID()
	if err != nil {
		return User{}, err
	}

	newUserName, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	return User{id: newUserID, name: newUserName}, nil
}

func (u *User) ChangeName(name string) error {
	newUserName, err := NewUserName(name)
	if err != nil {
		return err
	}

	u.name = newUserName

	return nil
}
