package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	user := r.Group("/user")
	user.POST("/login/")
	user.POST("/register/")

	if err := http.ListenAndServe(":8086", r); err != nil {
		log.Panic(err)
	}
}
