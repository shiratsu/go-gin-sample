package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func showIndexPage(ctx *gin.Context) {
	todos := dbGetAll()
	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
}

func createTodo(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbInsert(text, status)
	ctx.Redirect(302, "/")
}

func showDetail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := dbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})

}
func updateTodo(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

func deleteCheckTodo(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := dbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
}

func deleteTodo(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	dbDelete(id)
	ctx.Redirect(302, "/")
}
