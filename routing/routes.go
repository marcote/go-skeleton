package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/marcote/go-skeleton/controller"
)

var (
	//Router exposes the endpoint of the application.
	Router *gin.Engine
)

// Register API routes
func Register() {
	Router.GET("/ping", controller.Ping)
	Router.GET("/characters/:id", controller.RetrieveCharacters)
}

// Configure the routing behavior
func Configure() {
	Router = gin.New()
}
