package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ScannedItems struct {
	SKUs       map[string]int
	TotalPrice int
}

type Items struct {
	SKUs map[string]Prices
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
	var Checkout ScannedItems

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter SKUs in a comma-seperate line or exit to quit: ")

		scanner.Scan()

		text := scanner.Text()
		if text != "exit" {
			err := Checkout.Scan(text)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Please try again:")

			} else {
				var total int
				total, err = Checkout.GetTotalPrice()

				if err == nil {
					fmt.Println(fmt.Sprintf("Your total is: %d", total))
				} else {
					fmt.Println(err)
				}
			}

		} else {
			break
		}
	}
}

func (c *ScannedItems) Scan(SKU string) (err error) {

	(*c).SKUs = map[string]int{}
	SKUs := strings.Split(SKU, ",")

	for _, sku := range SKUs {

		if _, exists := Inventory.SKUs[sku]; !exists {
			err = errors.New(fmt.Sprintf("invalid sku: %s", sku))
			return
		}

		if val, exists := c.SKUs[sku]; exists {
			c.SKUs[sku] = val + 1
		} else {
			c.SKUs[sku] = 1
		}
	}

	return
}

func (c *ScannedItems) GetTotalPrice() (totalPrice int, err error) {

	for sku, quantity := range (*c).SKUs {

		if Inventory.SKUs[sku].Offers.Quantity != 0 && quantity >= Inventory.SKUs[sku].Offers.Quantity {

			totalPrice += Inventory.SKUs[sku].Offers.Price * (quantity / Inventory.SKUs[sku].Offers.Quantity)
			totalPrice += Inventory.SKUs[sku].DefaultPrice * (quantity % Inventory.SKUs[sku].Offers.Quantity)

		} else {
			totalPrice += Inventory.SKUs[sku].DefaultPrice * quantity
		}
	}

	return
}
