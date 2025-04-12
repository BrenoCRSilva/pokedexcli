package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandArea() Command {
	return Command{
		name:        "location",
		description: "",
		Callback:    commandArea,
	}
}

func commandArea(gs *game.GameState, param string) error {
	region, err := gs.PokeAPI.FetchLocation(param)
	if err != nil {
		return err
	}
	// TODO pagination logic
	for _, item := range region.Areas {
		fmt.Println(item.Name)
	}
	return nil
}
