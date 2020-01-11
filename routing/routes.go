package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/marcote/go-skeleton/controller"
)

var (
	Router *gin.Engine
)

// Register API routes
func Register() {
	Router.GET("/ping", controller.Ping)
	Router.GET("/characters/:id", controller.RetrieveCharacters)
}
