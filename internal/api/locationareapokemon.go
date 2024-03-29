package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreasPokemons(name string) (LocationAreaPokemonResponse, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + name

	cacheResp, cacheExists := cache.Get(fullURL)

	if cacheExists {
		result := LocationResultStruct{}
		er := json.Unmarshal(cacheResp, &result)

		if er != nil {
			fmt.Println(er)
		}
		return LocationAreaPokemonResponse{}, er
	}

	if len(name) <= 0 {
		return LocationAreaPokemonResponse{}, errors.New("length of location cannot be zero")
	}
	res, err := http.Get(fullURL)
	if res.StatusCode > 399 {
		fmt.Printf("Error status code : %v\n", res.StatusCode)
		return LocationAreaPokemonResponse{}, err
	}

	if err != nil {
		fmt.Println(err)
		return LocationAreaPokemonResponse{}, err

	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	result := LocationAreaPokemonResponse{}
	er := json.Unmarshal(body, &result)

	if er != nil {
		return LocationAreaPokemonResponse{}, er
	}

	return result, nil
}

type LocationAreaPokemonResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
