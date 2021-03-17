package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {

	r:=gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil{
		log.Fatal(err)
	}
}


