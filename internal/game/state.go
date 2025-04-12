package game

import (
	"time"

	"github.com/BrenoCRSilva/pokedexcli/internal/pokeapi"
)

type WorldPosition struct {
	Region   string
	Location string
	Area     string
}

type GameState struct {
	Mode             string // World, Pokedex, Inventory, Map or Encounter
	CurrentEncounter string
	PokeAPI          *pokeapi.Client
	CaughtPokemon    map[string]pokeapi.PokemonSpecies
	Position         WorldPosition
	LocationToRegion map[string]string
	Hometown         map[string]string
	Inventory        map[string]int
}

func NewGameState(cacheInterval time.Duration) *GameState {
	return &GameState{
		Mode:             "World",
		PokeAPI:          pokeapi.NewClient(cacheInterval),
		CaughtPokemon:    make(map[string]pokeapi.PokemonSpecies),
		Position:         WorldPosition{},
		LocationToRegion: make(map[string]string),
		Hometown: map[string]string{
			"kanto":  "pallet-town-area",
			"johto":  "new-bark-town-area",
			"hoenn":  "littleroot-town-area",
			"sinnoh": "twinleaf-town-area",
			"unova":  "nuvema-town-area",
			"kalos":  "vaniville-town-area",
			"alola":  "hauoli-outskirts-area",
			"galar":  "postwick-area",
			"paldea": "cabo-poco-area",
		},
	}
}
