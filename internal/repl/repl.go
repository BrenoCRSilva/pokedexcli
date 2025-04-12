package repl

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/BrenoCRSilva/pokedexcli/internal/commands"
	"github.com/BrenoCRSilva/pokedexcli/internal/game"
)

func StartRepl(input *bufio.Scanner, cr *commands.CommandRegistry, gs *game.GameState) {
	fmt.Println("Welcome to Pokemon, CLI version.")
	c, _ := cr.Get("help")
	c.Callback(gs, "")
	for {
		fmt.Printf("%s > ", gs.Mode)
		input.Scan()
		cleaned := cleanInput(input.Text())
		cmd := cleaned[0]
		var param string
		if len(cleaned) > 1 {
			param = cleaned[1]
		}
		if c, ok := cr.Get(cmd); !ok {
			fmt.Println("Unknown command")
		} else {
			c.Callback(gs, param)
		}
	}
}

func cleanInput(input string) []string {
	var cleaned []string
	for _, word := range strings.Fields(input) {
		word = strings.TrimSpace(word)
		if word == " " || word == "" {
			continue
		}
		word = strings.ToLower(word)
		cleaned = append(cleaned, word)
	}
	return cleaned
}
