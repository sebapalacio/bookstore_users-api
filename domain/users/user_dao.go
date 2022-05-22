package users

import (
	"fmt"
	"github.com/sebapalacio/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("The user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d already exist", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
