package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandLeave() Command {
	return Command{
		name:        "leave",
		description: "",
		Callback:    commandLeave,
	}
}

func commandLeave(gs *game.GameState, _ string) error {
	if gs.Mode == "World" {
		err := fmt.Errorf("Already at world")
		fmt.Println(err)
		return err
	}
	gs.Mode = "World"
	return nil
}
