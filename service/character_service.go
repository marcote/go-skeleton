package service

import (
	"errors"
	"path"
	"strconv"

	"github.com/marcote/go-skeleton/datasource"
	response "github.com/marcote/go-skeleton/response"
)

//ICharacterService provides API entrypoint and mocking capabilities
type ICharacterService interface {
	GetCharacterByID(characterID int64) (response.Character, error)
}

//CharacterService decoupling struct
type CharacterService struct {
}

//SWAPISource represents the SW API main interaction
var SWAPISource datasource.ISWAPISource

//GetCharacterByID returns a Person if we pass a valid non-zero ID
func (sw CharacterService) GetCharacterByID(characterID int64) (response.Character, error) {

	if characterID == 0 {
		return response.Character{}, errors.New("Character ID cannot be 0")
	}

	//This is where we coordinate the logic and the call to "different" APIs
	person, err := getSwapiSource().GetPersonByID(characterID)

	if err != nil {
		return response.Character{}, err
	}

	specieID, err := strconv.ParseInt(path.Base(string(person.SpeciesURLs[0])), 10, 32)
	spicie, err := getSwapiSource().GetSpicieByID(specieID)

	if err != nil {
		return response.Character{}, err
	}

	starshipID, err := strconv.ParseInt(path.Base(string(person.StarshipURLs[0])), 10, 32)
	starship, err := getSwapiSource().GetStarshipByID(starshipID)

	character := response.Character{
		Name:     person.Name,
		Height:   person.Height,
		Starship: starship.Name,
		Spicie:   spicie.Name,
	}

	return character, nil
}

func getSwapiSource() datasource.ISWAPISource {
	if SWAPISource == nil {
		swapiSource := datasource.SWAPISource{}
		return &swapiSource
	}
	return SWAPISource
}
