package users

import (
	"github.com/sebapalacio/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	//`json:"id"` is used in this example, for take the impute for the request, the formate json is minus
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestErr("invalid email address")
	}
	return nil
}
