package users

import (
	"github.com/sri-05/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user User) Validate() *errors.RestErr {
	mail := strings.TrimSpace(strings.ToLower(user.Email))
	if mail == "" {
		return errors.NewBadRequestError("invalid Email Address")
	}
	return nil

}
