package datasource

import (
	"errors"
	"fmt"

	r "github.com/marcote/go-skeleton/response"
	util "github.com/marcote/go-skeleton/util"
)

// ISWAPISource interface represents the contract of what we want to retrieve from SWAPI
type ISWAPISource interface {

	//GET People
	GetPersonByID(personID int64) (r.Person, error)

	//GET Spicie
	GetSpicieByID(spicieID int64) (r.Species, error)

	//GET Starships
	GetStarshipByID(starshipID int64) (r.Starship, error)

	//GET Homeworld
	GetHomeWorldByID(homeworldID int64) (r.Planet, error)
}

//HTTPClient is hour Client to retrieve information from SW API
var HTTPClient = util.NewClient()

//SWAPISource decoupling struct
type SWAPISource struct{}

//GetPersonByID using the SW API
func (ss SWAPISource) GetPersonByID(personID int64) (r.Person, error) {

	if personID == 0 {
		return r.Person{}, errors.New("personID ID cannot be 0")
	}

	req, err := HTTPClient.NewRequest(fmt.Sprintf("people/%d", personID))
	if err != nil {
		return r.Person{}, err
	}

	var person r.Person
	if _, err = HTTPClient.Do(req, &person); err != nil {
		return r.Person{}, err
	}

	return person, nil
}

//GetSpicieByID using the SW API
func (ss SWAPISource) GetSpicieByID(spicieID int64) (r.Species, error) {

	if spicieID == 0 {
		return r.Species{}, errors.New("spicieID ID cannot be 0")
	}

	req, err := HTTPClient.NewRequest(fmt.Sprintf("species/%d", spicieID))
	if err != nil {
		return r.Species{}, err
	}

	var spicie r.Species
	if _, err = HTTPClient.Do(req, &spicie); err != nil {
		return r.Species{}, err
	}

	return spicie, nil
}

//GetStarshipByID using the SW API
func (ss SWAPISource) GetStarshipByID(starshipID int64) (r.Starship, error) {

	if starshipID == 0 {
		return r.Starship{}, errors.New("starshipID ID cannot be 0")
	}

	req, err := HTTPClient.NewRequest(fmt.Sprintf("starships/%d", starshipID))
	if err != nil {
		return r.Starship{}, err
	}

	var starship r.Starship
	if _, err = HTTPClient.Do(req, &starship); err != nil {
		return r.Starship{}, err
	}

	return starship, nil
}

//GetHomeWorldByID using the SW API
func (ss SWAPISource) GetHomeWorldByID(homeworldID int64) (r.Planet, error) {

	if homeworldID == 0 {
		return r.Planet{}, errors.New("homeworldID ID cannot be 0")
	}

	req, err := HTTPClient.NewRequest(fmt.Sprintf("planets/%d", homeworldID))
	if err != nil {
		return r.Planet{}, err
	}

	var homeworld r.Planet
	if _, err = HTTPClient.Do(req, &homeworld); err != nil {
		return r.Planet{}, err
	}

	return homeworld, nil
}
