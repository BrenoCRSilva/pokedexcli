package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandExplore() Command {
	return Command{
		name:        "explore",
		description: "",
		Callback:    commandExplore,
	}
}

func commandExplore(gs *game.GameState, _ string) error {
	area, err := gs.PokeAPI.FetchLocationArea(gs.Position.Area)
	if err != nil {
		return err
	}
	encounterList := area.PokemonEncounters
	if len(encounterList) == 0 {
		err := fmt.Errorf("There are no pokemon here.")
		fmt.Println(err)
		return err
	}
	tileRate := 200
	encounterSeed := rand.UintN(uint(len(encounterList)))
	chance := encounterList[encounterSeed].VersionDetails[0].EncounterDetails[0].Chance
	for range 5 {
		game.ExploreAnimation()
		if rand.UintN(256) < uint(tileRate) {
			if rand.UintN(100) < uint(chance) {
				gs.Mode = "Encounter"
				fmt.Printf("A wild %s has appeared\n", encounterList[encounterSeed].Pokemon.Name)
				break
			}
		}
	}
	return nil
}
