package reponse

type personURL string
type filmURL string
type characterURL string
type speciesURL string
type vehicleURL string
type starshipURL string

// Character is an individual with spicie and starship
type Character struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Homeworld string `json:"homeworld"`
	Starship  string `json:"starship"`
	Spicie    string `json:"spicie"`
}

// Person is an individual person or character
type Person struct {
	Name         string        `json:"name"`
	Height       string        `json:"height"`
	Mass         string        `json:"mass"`
	HairColor    string        `json:"hair_color"`
	SkinColor    string        `json:"skin_color"`
	EyeColor     string        `json:"eye_color"`
	BirthYear    string        `json:"birth_year"`
	Gender       string        `json:"gender"`
	Homeworld    string        `json:"homeworld"`
	FilmURLs     []filmURL     `json:"films"`
	SpeciesURLs  []speciesURL  `json:"species"`
	VehicleURLs  []vehicleURL  `json:"vehicles"`
	StarshipURLs []starshipURL `json:"starships"`
	Created      string        `json:"created"`
	Edited       string        `json:"edited"`
	URL          string        `json:"url"`
}

// A Starship is a single transport craft that has hyperdrive capability.
type Starship struct {
	Name                 string      `json:"name"`
	Model                string      `json:"model"`
	Manufacturer         string      `json:"manufacturer"`
	CostInCredits        string      `json:"cost_in_credits"`
	Length               string      `json:"length"`
	MaxAtmospheringSpeed string      `json:"max_atmosphering_speed"`
	Crew                 string      `json:"crew"`
	Passengers           string      `json:"passengers"`
	CargoCapacity        string      `json:"cargo_capacity"`
	Consumables          string      `json:"consumables"`
	HyperdriveRating     string      `json:"hyperdrive_rating"`
	MGLT                 string      `json:"MGLT"`
	StarshipClass        string      `json:"starship_class"`
	PilotURLs            []personURL `json:"pilots"`
	FilmURLs             []filmURL   `json:"films"`
	Created              string      `json:"created"`
	Edited               string      `json:"edited"`
	URL                  string      `json:"url"`
}

// A Species is a type of person or character within the Star Wars Universe.
type Species struct {
	Name            string      `json:"name"`
	Classification  string      `json:"classification"`
	Designation     string      `json:"designation"`
	AverageHeight   string      `json:"average_height"`
	SkinColors      string      `json:"skin_colors"`
	HairColors      string      `json:"hair_colors"`
	EyeColors       string      `json:"eye_colors"`
	AverageLifespan string      `json:"average_lifespan"`
	Homeworld       string      `json:"homeworld"`
	Language        string      `json:"language"`
	PeopleURLs      []personURL `json:"people"`
	FilmURLs        []filmURL   `json:"films"`
	Created         string      `json:"created"`
	Edited          string      `json:"edited"`
	URL             string      `json:"url"`
}
