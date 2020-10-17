package adhandler

import (
	"github.com/checkPrice/pkg/parser"
	log "github.com/sirupsen/logrus"
)


type Ad struct{
	Id			int
	Number		string
	Price		string
	NewPrice 	string
	Email 		string
}

func (a *Ad) IsPriceChange() bool {
	var err error
	a.NewPrice, err = parser.GetPrice(a.Number)
	if err != nil{
		log.Error(err)
	}
	if a.NewPrice != a.Price && a.NewPrice != "" {
		return true
	}
	return false
}