package main

import (
	"github.com/antony-raul/tw/config"
	"github.com/antony-raul/tw/handlers/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	r := gin.New()

	api := r.Group("api")

	user.Router(api.Group("user"))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
