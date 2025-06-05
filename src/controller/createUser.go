package controller

import (
	"fmt"
	
	"github.com/lricardogarcia/primeiro-crud-go/src/configuration/rest_err"	
	"github.com/lricardogarcia/primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println("User created successfully:", userRequest)

}