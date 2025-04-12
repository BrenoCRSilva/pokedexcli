package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandCaught() Command {
	return Command{
		name:        "caught",
		description: "",
		Callback:    commandCaught,
	}
}

func commandCaught(gs *game.GameState, _ string) error {
	if gs.Mode != "Pokedex" {
		err := fmt.Errorf("Invalid command")
		fmt.Println(err)
		return err
	}
	fmt.Println("Your caught pokemon:")
	for pokemon, species := range gs.CaughtPokemon {
		var entry int
		for _, r := range species.PokedexNumbers {
			if r.Pokedex.Name == "national" {
				entry = r.EntryNumber
				break
			}
		}
		fmt.Printf("%d : %s\n", entry, pokemon)
	}
	return nil
}
