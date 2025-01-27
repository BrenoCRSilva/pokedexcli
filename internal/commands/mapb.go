package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandMapb() Command {
	return Command{
		name:        "mapb",
		description: "Displays previous 20 location areas",
		Callback:    commandMapb,
	}
}

func commandMapb(gs *game.GameState, _ string) error {
	locationAreas, err := gs.PokeAPI.ListPreviousLocationAreas()
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
