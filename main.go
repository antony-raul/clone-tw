package main

import (
	"github.com/antony-raul/tw/handlers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	api := r.Group("api")

	user.Router(api.Group("user"))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
