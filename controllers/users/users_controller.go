package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sri-05/bookstore_users-api/domain/users"
	"github.com/sri-05/bookstore_users-api/services"
	"github.com/sri-05/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("invalid json body")

		c.JSON(restError.Status, restError)
		return
	}
	fmt.Println("Received input data is after unmarshall", user)
	result, SaveErr := services.CreateUser(user)
	if SaveErr != nil {
		c.JSON(SaveErr.Status, SaveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userid, userErr := strconv.ParseInt(c.Param("user_id"), 10, 14)
	if userErr != nil {
		fmt.Println("Found error while getting query parameter")
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
	}
	result, getErr := services.GetUser(userid)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
