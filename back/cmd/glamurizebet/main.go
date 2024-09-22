package main

import (
	appcontext "glamurizebet/internal/application"
	"glamurizebet/internal/endpoints"
	endpoint_util "glamurizebet/internal/endpoints/util"

	"github.com/gin-gonic/gin"
)

func main() {

	env, err := appcontext.LoadEnv()
	if err != nil {
		panic(err)
	}

	appcontext, err := appcontext.GetAppContext(&env)
	if err != nil {
		panic(err)
	}

	appcontext.Log.Info.Println("configuring GIN mode")
	if appcontext.IsProd {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		gin.DefaultWriter = appcontext.Log.Info.Writer()

		appcontext.Log.Info.Println("gin setted to release mode")
	} else {
		gin.SetMode(gin.DebugMode)
		appcontext.Log.Info.Println("gin setted to debug mode")
	}

	router := gin.Default()

	appcontext.Log.Info.Println("creating routes")
	v1 := router.Group("/api/v1")
	{
		external := v1.Group("/external")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.EXTERNAL, external, &appcontext)
		}

		play := v1.Group("/play")
		{
			endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.PLAY, play, &appcontext)
		}

		endpoints.AddRoutesBy(endpoint_util.V1, endpoint_util.COMMON, v1, &appcontext)
	}

	appcontext.Log.Info.Println("starting server")
	router.Run("localhost:8000")
}
