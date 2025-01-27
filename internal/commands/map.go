package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandMap() Command {
	return Command{
		name:        "map",
		description: "Displays next 20 location areas",
		Callback:    commandMap,
	}
}

func commandMap(gs *game.GameState, _ string) error {
	locationAreas, err := gs.PokeAPI.ListNextLocationAreas()
	if err != nil {
		return err
	}
	for _, item := range locationAreas.Results {
		fmt.Println(item.Name)
	}
	gs.PokeAPI.PageConfig.Next = locationAreas.Next
	gs.PokeAPI.PageConfig.Previous = locationAreas.Previous

	return nil
}
