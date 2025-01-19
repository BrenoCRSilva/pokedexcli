package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}
type LocationArea struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type APIResponse struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []LocationArea
}

type Config struct {
	Previous *string
	Next     *string
}

func main() {
	cfg := &Config{
		Previous: nil,
		Next:     nil,
	}
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
		},
	}
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		lower := strings.ToLower(input.Text())
		split := strings.Fields(lower)
		cmd := split[0]
		if c, ok := commands[cmd]; !ok {
			fmt.Println("Unknown command")
		} else if cmd == "help" {
			c.callback(cfg)
			for key := range commands {
				fmt.Println(key, ":", commands[key].description)
			}
		} else {
			c.callback(cfg)
		}

	}
}

func cleanInput(input string) []string {
	var cleaned []string
	split := strings.Split(input, " ")
	for _, word := range split {
		var tmp string
		tmp = strings.TrimSpace(word)
		if word == " " || word == "" {
			continue
		}
		tmp = strings.ToLower(tmp)
		cleaned = append(cleaned, tmp)
	}
	return cleaned
}

func commandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	return nil
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return err
	}
	for _, a := range apiResponse.Results {
		fmt.Println(a.Name)
	}

	cfg.Next = apiResponse.Next
	return nil
}

func commandMapb(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil {
		url = *cfg.Previous
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return err
	}
	for _, a := range apiResponse.Results {
		fmt.Println(a.Name)
	}

	cfg.Previous = apiResponse.Previous
	return nil
}
