package services

import (
	"fmt"
	"github.com/sri-05/bookstore_users-api/domain/users"
	"github.com/sri-05/bookstore_users-api/utils/errors"
	"reflect"
)

func GetUser(userid int64) (*users.User, *errors.RestErr) {
	fmt.Println("type of userid in getuser func is", reflect.TypeOf(userid))
	fmt.Println(userid)
	result := &users.User{Id: userid}
	fmt.Println("result", result)
	er := result.Get()
	if er != nil {
		return nil, er
	}
	fmt.Println("Type of userid after conv  is", reflect.TypeOf(result))
	fmt.Println(result)
	return result, nil

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	er := user.Validate()
	if er != nil {
		return nil, er
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
