package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Name         string         `json:"name"`
	LineStatuses []LineStatuses `json:"lineStatuses"`
}

type LineStatuses struct {
	Description string `json:"statusSeverityDescription"`
}

func main() {
	var respStructs []Response
	response, err := http.Get("https://api.tfl.gov.uk/Line/central/Status")
	if err != nil {
		log.Fatalln(err)
	}

	var rawJSON []json.RawMessage

	data, err := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(data, &rawJSON); err != nil {
		log.Fatalln(err)
	}

	for _, object := range rawJSON {
		demarsResp := Response{}
		if err := json.Unmarshal(object, &demarsResp); err != nil {
			log.Fatalln(err)
		}
		respStructs = append(respStructs, demarsResp)
	}
	fmt.Println(respStructs[0].LineStatuses[0].Description)
}
