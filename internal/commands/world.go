package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandWorld() Command {
	return Command{
		name:        "world",
		description: "",
		Callback:    commandWorld,
	}
}

func commandWorld(gs *game.GameState, _ string) error {
	regions, err := gs.PokeAPI.ListRegions()
	if err != nil {
		return err
	}
	for _, item := range regions.Results {
		fmt.Println(item.Name)
	}
	return nil
}
