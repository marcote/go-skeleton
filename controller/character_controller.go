package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcote/go-skeleton/service"
)

//CharacterService represents the SW API main interaction
var CharacterService service.ICharacterService

// Ping response is a Pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//RetrieveCharacters with specie and planets
// GET /characters/:id
// example: GET /characters/1
func RetrieveCharacters(c *gin.Context) {

	//parameter validation
	characterID, errParsing := strconv.ParseInt(c.Param("id"), 10, 64)

	if errParsing != nil {
		log.Println("The request is badly formated")
		c.JSON(http.StatusBadRequest, errParsing)
		return
	}

	//call service here.
	response, errService := getCharacterService().GetCharacterByID(characterID)

	if errService != nil {
		log.Println("Error calling service")
		c.JSON(http.StatusInternalServerError, errService)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func getCharacterService() service.ICharacterService {
	if CharacterService == nil {
		apiService := service.CharacterService{}
		return &apiService
	}
	return CharacterService
}
