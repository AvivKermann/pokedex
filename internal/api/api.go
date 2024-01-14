package api

import(
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	
)

func GetLocationAreas(defultURL string)  (LocationResultStruct, error) {

res, err := http.Get(defultURL)
if res.StatusCode > 399 {
	fmt.Printf("Error status code : %v\n", res.StatusCode)
	return LocationResultStruct{}, err
}

if err != nil {
	fmt.Println(err)
	return LocationResultStruct{}, err

}

defer res.Body.Close()
body, _ := io.ReadAll(res.Body)
result := LocationResultStruct{} 
er := json.Unmarshal(body, &result)

if er != nil {
	fmt.Println(er)
	return LocationResultStruct{}, err
}

return result, nil
}

type LocationResultStruct struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}