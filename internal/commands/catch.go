package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandCatch() Command {
	return Command{
		name:        "catch",
		description: "Attempts to catch the pokemon given as parameter",
		Callback:    commandCatch,
	}
}

func commandCatch(gs *game.GameState, _ string) error {
	if gs.Mode != "Encounter" {
		err := fmt.Errorf("There's a time and place for everything.")
		fmt.Println(err)
		return err
	}
	pokemon, err := gs.PokeAPI.FetchPokemonSpecies(gs.CurrentEncounter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for range 3 {
		game.ShakeAnimation()
	}
	catchRate := pokemon.CaptureRate
	if rand.IntN(100) < catchRate*100/255 {
		fmt.Printf("Congratulations, you've caught %s.\n", pokemon.Name)
		gs.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println("Aah, almost had it.")
	}
	return nil
}
