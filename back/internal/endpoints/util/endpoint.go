package endpoint_util

import (
	appcontext "glamurizebet/internal/application"

	"github.com/gin-gonic/gin"
)

type Versao uint8

const (
	V1 Versao = iota
)

type Type uint8

const (
	EXTERNAL Type = iota
	COMMON
	PLAY
)

type Endpoint interface {
	AddRoutes(router *gin.RouterGroup, appContext *appcontext.AppContext)
	GetVersion() Versao
	GetType() Type
}
