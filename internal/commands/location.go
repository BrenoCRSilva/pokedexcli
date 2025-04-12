package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandLocation() Command {
	return Command{
		name:        "location",
		description: "",
		Callback:    commandLocation,
	}
}

func commandLocation(gs *game.GameState, param string) error {
	region, err := gs.PokeAPI.FetchRegion(param)
	if err != nil {
		return err
	}
	// TODO pagination logic
	for _, item := range region.Locations {
		fmt.Println(item.Name)
	}
	return nil
}
