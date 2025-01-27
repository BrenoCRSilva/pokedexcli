package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandCatch() Command {
	return Command{
		name:        "catch",
		description: "Attempts to catch the pokemon given as parameter",
		Callback:    commandCatch,
	}
}

func commandCatch(gs *game.GameState, param string) error {
	pokemon, err := gs.PokeAPI.FetchPokemon(param)
	if err != nil {
		return err
	}
	fmt.Println(pokemon)
	// TODO catching logic
	return nil
}
