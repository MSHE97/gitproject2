package card

import "github.com/MSHE97/gitproject2/pkg/types"

func Withdraw(card types.Card) types.Card {
	card.Balance -= 100
	return card
}
