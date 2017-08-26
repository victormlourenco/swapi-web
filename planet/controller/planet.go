package controller

import (
	"net/http"

	planetApi "github.com/victormlourenco/swapi-web/planet"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var r *render.Render

// Create : Create HTTP routes
func Create(router *gin.RouterGroup) {
	planet := router.Group("")
	{
		planet.GET("/", getPlanet)
	}

	r = render.New(render.Options{
		IndentJSON: true,
		Directory:  "planet/template",
		Extensions: []string{".html"},
	})
}

func getPlanet(c *gin.Context) {
	planet, err := planetApi.PickPlanet()
	if err != nil {
		log.WithFields(log.Fields{
			"Host":    c.Request.Host,
			"Request": c.Request.RequestURI,
		}).Error(err)
		r.HTML(c.Writer, http.StatusBadRequest, "planet", planet)
		return
	}
	r.HTML(c.Writer, http.StatusOK, "planet", planet)
}
