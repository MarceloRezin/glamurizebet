package endpoint_util

import (
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
	AddRoutes(router *gin.RouterGroup)
	GetVersion() Versao
	GetType() Type
}
