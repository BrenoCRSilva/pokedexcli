package commands

import (
	"fmt"
	"os"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandExit() Command {
	return Command{
		name:        "exit",
		description: "Exits the pokedex.",
		Callback:    commandExit,
	}
}

func commandExit(_ *game.GameState, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
