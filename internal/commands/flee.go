package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandFlee() Command {
	return Command{
		name:        "flee",
		description: "",
		Callback:    commandFlee,
	}
}

func commandFlee(gs *game.GameState, _ string) error {
	if gs.Mode != "Encounter" {
		err := fmt.Errorf("Invalid command")
		fmt.Println(err)
		return err
	}
	chance := 20
	if rand.UintN(100) >= uint(chance) {
		fmt.Println("Got away safely.")
		gs.Mode = "World"
	} else {
		fmt.Println("Can't escape!")
	}
	return nil
}
