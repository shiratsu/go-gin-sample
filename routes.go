// routes.go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	if router == nil {
		fmt.Println("router null") // nil だったら数字を入れてnを使いたい
	}

	// Handle the index route
	//Index
	router.GET("/", showIndexPage)

	//Create
	router.POST("/new", createTodo)

	//Detail
	router.GET("/detail/:id", showDetail)

	//Update
	router.POST("/update/:id", updateTodo)

	//削除確認
	router.GET("/delete_check/:id", deleteCheckTodo)

	//Delete
	router.POST("/delete/:id", deleteTodo)

	// Start serving the application
	dbInit()

	router.Run()

}
