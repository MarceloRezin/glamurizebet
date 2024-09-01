package main

import (
	"glamurizebet/internal/endpoints"
	endpoint_util "glamurizebet/internal/endpoints/util"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		external := v1.Group("/external")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.EXTERNAL, external)
		}

		play := v1.Group("/play")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.PLAY, play)
		}

		endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.COMMON, v1)
	}

	router.Run("localhost:8000")
}
