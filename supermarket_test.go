package main

import (
	"testing"
)

type checkout struct {
	SKUs          string
	totalExpected int
}

var addTests = []checkout{
	checkout{"B,A,B", 95},
	checkout{"A,B", 80},
	checkout{"C,D", 35},
}

func Test(t *testing.T) {
	for _, test := range addTests {
		t.Run(test.SKUs, func(t *testing.T) {

			var Checkout ScannedItems
			_ = Checkout.Scan(test.SKUs)

			total, _ := Checkout.GetTotalPrice()

			if total != test.totalExpected {
				t.Errorf("Output %d not equal to expected %d", total, test.totalExpected)
			}
		})
	}
}
