package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Items struct {
	SKU map[string]Prices
}

type Prices struct {
	DefaultPrice int
	Offers       struct {
		Quantity int
		Price    int
	}
}

var Inventory Items

func init() {

	jsonFile, err := os.Open("inventory.json")

	if err == nil {
		var fileContents []byte
		fileContents, err := ioutil.ReadAll(jsonFile)

		if err == nil {
			json.Unmarshal([]byte(fileContents), &Inventory)
			jsonFile.Close()
		}
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		os.Exit(1)
	}

}

func main() {

}
