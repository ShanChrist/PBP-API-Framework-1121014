package main

import (
	"log"

	"github.com/ExplorasiGIN/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/insertUser", controllers.InsertUser)
	router.PUT("/updateUser/:id", controllers.UpdateUser)
	router.DELETE("/deleteUser/:id", controllers.DeleteUser)
	router.GET("/user/:id", controllers.GetSpecificUser)

	err := router.Run(":7777")
	if err != nil {
		log.Fatal(err)
	}
}
