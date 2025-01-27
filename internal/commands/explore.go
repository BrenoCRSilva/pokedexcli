package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandExplore() Command {
	return Command{
		name:        "explore",
		description: "",
		Callback:    commandExplore,
	}
}

func commandExplore(gs *game.GameState, param string) error {
	locationArea, err := gs.PokeAPI.FetchLocationArea(param)
	if err != nil {
		return err
	}
	for _, encounters := range locationArea.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil
}
