package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandMove() Command {
	return Command{
		name:        "move",
		description: "",
		Callback:    commandMove,
	}
}

func commandMove(gs *game.GameState, param string) error {
	locationArea, err := gs.PokeAPI.FetchLocationArea(param)
	if err != nil {
		return err
	}
	if gs.Position.Region == "" {
		err = fmt.Errorf("You must fly to a region first.")
		fmt.Println(err)
		return err
	}
	area := locationArea.Name
	location := locationArea.Location
	region := gs.LocationToRegion[location.Name]
	if gs.LocationToRegion[location.Name] != gs.Position.Region {
		err = fmt.Errorf("Area out of bounds.")
		fmt.Println(err)
		return err
	}
	gs.Position = game.WorldPosition{
		Region:   region,
		Location: location.Name,
		Area:     area,
	}
	fmt.Printf("You are now in %s, %s: %s\n", region, location.Name, area)
	return nil
}
