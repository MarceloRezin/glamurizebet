package login

import (
	appcontext "glamurizebet/internal/application"
	endpoint_util "glamurizebet/internal/endpoints/util"

	"github.com/gin-gonic/gin"
)

type LoginEndpoint struct{}

func (LoginEndpoint) AddRoutes(router *gin.RouterGroup, appContext *appcontext.AppContext) {
	lb := login_business{appContext: appContext}

	router.GET("/login", lb.logar)
	router.POST("/logar", lb.logar)
}

func (LoginEndpoint) GetVersion() endpoint_util.Versao {
	return endpoint_util.V1
}

func (LoginEndpoint) GetType() endpoint_util.Type {
	return endpoint_util.EXTERNAL
}

/*
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func logar(c *gin.Context) {
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	c.IndentedJSON(http.StatusOK, albums)
}
*/
