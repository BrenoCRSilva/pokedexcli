package commands

import (
	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandPokedex() Command {
	return Command{
		name:        "pokedex",
		description: "Displays all your caught pokemon",
		Callback:    commandPokedex,
	}
}

func commandPokedex(gs *game.GameState, _ string) error {
	gs.Mode = "Pokedex"
	return nil
}
