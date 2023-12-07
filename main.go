package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pwd, _ := os.Getwd()
	r.LoadHTMLGlob(pwd + "/views/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/send-information", func(ctx *gin.Context) {
		inputValue := ctx.PostForm("information")
		result := "Processed: " + inputValue

		ctx.HTML(http.StatusOK, "results.html", gin.H{
			"result": result,
		})
	})

	r.Run(":8081")
}
