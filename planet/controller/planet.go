package controller

import (
	"net/http"

	planetApi "github.com/victormlourenco/swapi-web/planet"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

var r *render.Render

// Create : Create HTTP routes
func Create(router *gin.RouterGroup) {
	equipe := router.Group("")
	{
		equipe.GET("/", getPlanet)
	}

	r = render.New(render.Options{
		IndentJSON: true,
		Directory:  "planet/template",
		Extensions: []string{".html"},
	})
}

func getPlanet(c *gin.Context) {
	planet, _ := planetApi.PickPlanet()
	r.HTML(c.Writer, http.StatusOK, "planet", planet)
}
