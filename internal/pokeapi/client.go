package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BrenoCRSilva/pokedexcli/internal/cache"
)

type Client struct {
	baseUrl    string
	cache      *cache.Cache
	client     *http.Client
	PageConfig PageConfig
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		baseUrl: "https://pokeapi.co/api/v2",
		cache:   cache.NewCache(cacheInterval),
		client:  &http.Client{},
	}
}

func (c *Client) fetchFromCacheOrAPI(url string) ([]byte, error) {
	if cached, ok := c.cache.Get(url); ok {
		return cached, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	c.cache.Add(url, data)
	return data, nil
}

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	url := fmt.Sprintf("%s/pokemon/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func (c *Client) FetchPokemonSpecies(name string) (PokemonSpecies, error) {
	var pokemon PokemonSpecies
	url := fmt.Sprintf("%s/pokemon-species/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return PokemonSpecies{}, err
	}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return pokemon, nil
}

func (c *Client) ListNextLocationAreas() (PageResponse[LocationArea], error) {
	var locationAreas PageResponse[LocationArea]
	url := fmt.Sprintf("%s/location-area/", c.baseUrl)
	if c.PageConfig.Next != nil {
		url = *c.PageConfig.Next
	}
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return PageResponse[LocationArea]{}, err
	}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return PageResponse[LocationArea]{}, err
	}
	return locationAreas, nil
}

func (c *Client) ListPreviousLocationAreas() (PageResponse[LocationArea], error) {
	var locationAreas PageResponse[LocationArea]
	url := fmt.Sprintf("%s/location-area/", c.baseUrl)
	if c.PageConfig.Previous != nil {
		url = *c.PageConfig.Previous
	}
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return PageResponse[LocationArea]{}, err
	}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return PageResponse[LocationArea]{}, err
	}
	return locationAreas, nil
}

func (c *Client) FetchLocationArea(name string) (LocationArea, error) {
	var locationArea LocationArea
	url := fmt.Sprintf("%s/location-area/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return LocationArea{}, err
	}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (c *Client) ListRegions() (PageResponse[Region], error) {
	var regions PageResponse[Region]
	// regions don't have need for page navigation
	url := fmt.Sprintf("%s/region/", c.baseUrl)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return PageResponse[Region]{}, err
	}
	err = json.Unmarshal(data, &regions)
	if err != nil {
		return PageResponse[Region]{}, err
	}
	return regions, nil
}

func (c *Client) FetchRegion(name string) (Region, error) {
	var region Region
	url := fmt.Sprintf("%s/region/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return Region{}, err
	}
	err = json.Unmarshal(data, &region)
	if err != nil {
		return Region{}, err
	}
	return region, nil
}

func (c *Client) FetchLocation(name string) (Location, error) {
	var location Location
	url := fmt.Sprintf("%s/location/%s", c.baseUrl, name)
	data, err := c.fetchFromCacheOrAPI(url)
	if err != nil {
		return Location{}, err
	}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}
