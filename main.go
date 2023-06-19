package main

import (
	"boilerplate/config"
	post_user "boilerplate/controller/register/post"
	sql_connection "boilerplate/library/mysql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	configData := config.LoadEnv()

	//localization.LoadBundle(configData.Server)
	sqlErr := sql_connection.NewConnection(configData.MySQL)
	if sqlErr != nil {
		panic(sqlErr)
	}

	fmt.Println("starting.............")
	router := gin.Default()
	fmt.Println(router)
	router.POST("/Registers", post_user.PostUserInfo)
	router
	fmt.Println("post call")
	//router.POST("/login", post_user.LoginHandler)
	router.Run("localhost:8080")
}
