package card

import (
	"github.com/MSHE97/gitproject2/pkg/types"
)

func Issue() types.Card {
	return types.Card{
		Balance:  0,
		Currency: types.Currency("TJS"),
	}
}
