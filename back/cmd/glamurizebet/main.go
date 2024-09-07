package main

import (
	appcontext "glamurizebet/internal/application"
	"glamurizebet/internal/db"
	"glamurizebet/internal/endpoints"
	endpoint_util "glamurizebet/internal/endpoints/util"

	"github.com/gin-gonic/gin"
)

func main() {

	env, err := appcontext.LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := db.Create(&env)

	if err != nil {
		panic("failed to connect database")
	}

	ac := appcontext.AppContext{
		DB: db,
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		external := v1.Group("/external")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.EXTERNAL, external, &ac)
		}

		play := v1.Group("/play")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.PLAY, play, &ac)
		}

		endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.COMMON, v1, &ac)
	}

	router.Run("localhost:8000")
}
