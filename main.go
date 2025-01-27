package main

import (
	"bufio"
	"os"
	"time"

	"github.com/BrenoCRSilva/pokedexcli/internal/commands"
	"github.com/BrenoCRSilva/pokedexcli/internal/game"
	"github.com/BrenoCRSilva/pokedexcli/internal/repl"
)

func main() {
	registry := commands.NewCommandRegistry()
	input := bufio.NewScanner(os.Stdin)
	state := game.NewGameState(5 * time.Minute)
	repl.StartRepl(input, registry, state)
}
