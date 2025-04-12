package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandWhereami() Command {
	return Command{
		name:        "whereami",
		description: "",
		Callback:    commandWhereami,
	}
}

func commandWhereami(gs *game.GameState, _ string) error {
	fmt.Printf(
		"You are  in %s, %s: %s\n",
		gs.Position.Region,
		gs.Position.Location,
		gs.Position.Area,
	)
	return nil
}
