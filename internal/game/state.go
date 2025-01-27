package game

import (
	"time"

	"github.com/BrenoCRSilva/pokedexcli/internal/pokeapi"
)

type GameState struct {
	PokeAPI       *pokeapi.Client
	CaughtPokemon map[string]pokeapi.Pokemon
}

func NewGameState(cacheInterval time.Duration) *GameState {
	return &GameState{
		PokeAPI:       pokeapi.NewClient(cacheInterval),
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
}
