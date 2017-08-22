package main

import (
	"net/http"
	"swapi-web/config"
	planetCtrl "swapi-web/planet/controller"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	publicRoute := router.Group("/")
	planetCtrl.Create(publicRoute)
	s := &http.Server{
		Addr:    config.GetString("ServerAddress"),
		Handler: router,
	}

	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
