package users

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sri-05/bookstore_users-api/datasources/mysql/users_db"
	"github.com/sri-05/bookstore_users-api/utils/date_utils"
	"github.com/sri-05/bookstore_users-api/utils/errors"
)

var (
	UserDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := UserDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d is not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil

}

func (user *User) Save() *errors.RestErr {
	err := users_db.Client.Ping()
	if err != nil {
		panic(err)
	}
	result := UserDB[user.Id]
	if result != nil {
		if result.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user email %s is already available", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d is already avaiable", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()

	UserDB[user.Id] = user

	return nil

}
