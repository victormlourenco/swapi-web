package main

import (
	"net/http"

	"github.com/victormlourenco/swapi-web/config"
	planetCtrl "github.com/victormlourenco/swapi-web/planet/controller"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	s := &http.Server{
		Addr:    config.GetString("ServerAddress"),
		Handler: getMainEngine(),
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func getMainEngine() *gin.Engine {
	router := gin.Default()
	publicRoute := router.Group("/")
	planetCtrl.Create(publicRoute)
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))
	return router
}
