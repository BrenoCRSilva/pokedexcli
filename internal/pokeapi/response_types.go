package pokeapi

type PageResponse[T any] struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []T     `json:"results"`
}

type PageConfig struct {
	Previous *string
	Next     *string
}

type Region struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Locations []Location `json:"locations"`
}

type Location struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Areas []LocationArea `json:"areas"`
}

type LocationArea struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	Location          Location           `json:"location"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type VersionEncounterDetail struct {
	EncounterDetails []Encounter `json:"encounter_details"`
}

type Encounter struct {
	Chance int `json:"chance"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon                  `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

type PokemonSpecies struct {
	Name           string                   `json:"name"`
	CaptureRate    int                      `json:"capture_rate"`
	PokedexNumbers []PokemonSpeciesDexExtry `json:"pokedex_numbers"`
}

type PokemonSpeciesDexExtry struct {
	EntryNumber int     `json:"entry_number"`
	Pokedex     Pokedex `json:"pokedex"`
}

type Pokedex struct {
	Name string `json:"name"`
}
