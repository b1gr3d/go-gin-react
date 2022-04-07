package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// these are the endpoints
	//C
	router.POST("/change/create", routes.AddChange)
	//R
	router.GET("/env/:env", routes.GetChangesByEnv)
	router.GET("/changes", routes.GetChanges)
	router.GET("/change/:id/", routes.GetChangeById)
	//U
	router.PUT("/env/update/:id", routes.UpdateEnv)
	router.PUT("/change/update/:id", routes.UpdateChange)
	//D
	router.DELETE("/change/delete/:id", routes.DeleteChange)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
