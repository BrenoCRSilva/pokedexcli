package commands

import (
	"fmt"
	"strings"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func NewCommandFly() Command {
	return Command{
		name:        "fly",
		description: "",
		Callback:    commandFly,
	}
}

func commandFly(gs *game.GameState, param string) error {
	region, err := gs.PokeAPI.FetchRegion(param)
	if err != nil {
		return err
	}
	gs.Position.Region = param
	gs.Position.Location = strings.TrimSuffix(gs.Hometown[param], "-area")
	gs.Position.Area = gs.Hometown[param]
	for _, item := range region.Locations {
		location := item.Name
		gs.LocationToRegion[location] = param
	}
	fmt.Printf(
		"You are now in %s, %s: %s\n",
		gs.Position.Region,
		gs.Position.Location,
		gs.Position.Area,
	)
	return nil
}
