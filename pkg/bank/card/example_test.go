package card

import (
	"bank/pkg/bank/types"
	"fmt"
	
)
func ExampleTotal()  {
	cards := []types.Card {
		{
		Balance : 1_000_00,
		Active : true,
	},
}
fmt.Println(Total(cards))
// Output: 100000
}

func ExamplePaymentSources()  {
	card := []types.Card{
		{
			Balance: 1_000_00,
			PAN: "7777 xxxx xxxx 7777",
			Active: true,
		},
	}
	for _, Cards := range PaymentSources(card) {	
		fmt.Println(Cards.Number)	
	}
	// Output: 7777 xxxx xxxx 7777

}
func ExamplePaymentSources_inActive()  {
	card := []types.Card{
		{
			Balance: 1_000_00,
			PAN: "7777 xxxx xxxx 7777",
			Active: false,
		},
	}
	for _, Cards := range PaymentSources(card) {	
		fmt.Println(Cards.Number)	
	}
	// Output: 
}
func ExamplePaymentSources_noBalance()  {
	card := []types.Card{
		{
			Balance: 0,
			PAN: "7777 xxxx xxxx 7777",
			Active: true,
		},
	}
	for _, Cards := range PaymentSources(card) {	
		fmt.Println(Cards.Number)	
	}
	// Output:
}