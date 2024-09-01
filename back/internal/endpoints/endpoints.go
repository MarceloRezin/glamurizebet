package endpoints

import (
	"glamurizebet/internal/endpoints/login"
	endpoint_util "glamurizebet/internal/endpoints/util"

	"github.com/gin-gonic/gin"
)

var endpoints = []endpoint_util.Endpoint{
	login.LoginEndpoint{},
}

func getBy(version endpoint_util.Versao, tipe endpoint_util.Type) []endpoint_util.Endpoint {
	var selecionados = []endpoint_util.Endpoint{}

	for _, e := range endpoints {
		if e.GetVersion() == version && e.GetType() == tipe {
			selecionados = append(selecionados, e)
		}
	}

	return selecionados
}

func AddRoutesBy(versao endpoint_util.Versao, tipe endpoint_util.Type, router *gin.RouterGroup) {
	edps := getBy(versao, tipe)

	for _, e := range edps {
		e.AddRoutes(router)
	}
}
