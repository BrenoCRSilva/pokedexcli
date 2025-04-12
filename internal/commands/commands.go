package commands

import (
	"fmt"

	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

type Command struct {
	name        string
	description string
	Callback    func(*game.GameState, string) error
}

type CommandRegistry struct {
	registry map[string]Command
}

func (cr *CommandRegistry) Add(name string, cmd Command) {
	cr.registry[name] = cmd
}

func (cr *CommandRegistry) Get(name string) (Command, bool) {
	cmd, ok := cr.registry[name]
	return cmd, ok
}

func NewCommandRegistry() *CommandRegistry {
	registry := &CommandRegistry{
		registry: make(map[string]Command),
	}
	registry.Add("help", Command{
		name:        "help",
		description: "",
		Callback: func(*game.GameState, string) error {
			for name, cmd := range registry.registry {
				fmt.Printf("- %s: %s\n", name, cmd.description)
			}
			return nil
		},
	})
	registry.Add("map", NewCommandMap())
	registry.Add("mapb", NewCommandMapb())
	registry.Add("move", NewCommandMove())
	registry.Add("catch", NewCommandCatch())
	registry.Add("exit", NewCommandExit())
	registry.Add("pokedex", NewCommandPokedex())
	registry.Add("world", NewCommandWorld())
	registry.Add("location", NewCommandLocation())
	registry.Add("fly", NewCommandFly())
	registry.Add("area", NewCommandArea())
	registry.Add("caught", NewCommandCaught())
	registry.Add("leave", NewCommandLeave())
	registry.Add("explore", NewCommandExplore())
	registry.Add("whereami", NewCommandWhereami())
	registry.Add("flee", NewCommandFlee())
	return registry
}
