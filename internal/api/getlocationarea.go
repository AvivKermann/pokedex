package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AvivKermann/pokedex/internal/pokecache"
)

const interval = time.Minute * 5

var cache pokecache.Cache = pokecache.NewCache(interval)

func GetLocationAreas(defultURL string) (LocationResultStruct, error) {

	cacheResp, cacheExists := cache.Get(defultURL)

	if cacheExists {
		result := LocationResultStruct{}
		er := json.Unmarshal(cacheResp, &result)

		if er != nil {
			fmt.Println(er)
		}

		return result, nil
	}

	res, err := http.Get(defultURL)
	if res.StatusCode > 399 {
		fmt.Printf("Error status code : %v\n", res.StatusCode)
		return LocationResultStruct{}, err
	}

	if err != nil {
		fmt.Println(err)
		return LocationResultStruct{}, err

	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	cacheError := cache.Add(defultURL, body)
	if cacheError != nil {
		fmt.Println(cacheError)
	}

	result := LocationResultStruct{}
	er := json.Unmarshal(body, &result)

	if er != nil {
		fmt.Println(er)
		return LocationResultStruct{}, err
	}

	return result, nil
}

type LocationResultStruct struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
