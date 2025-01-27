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

type LocationArea struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
