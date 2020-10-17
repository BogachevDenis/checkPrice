
package adhandler

import (
	"github.com/checkPrice/pkg/parser"
	"testing"
)

func TestIsPriceChange(t *testing.T)  {

	type testpair struct {
		value 	string
		want 	bool
	}
	var number string
	number = "1892220447"
	price, _ := parser.GetPrice(number)

	var tests = []testpair{
		{ price, false },
		{ "2",	true },
	}
	for _, pair := range tests {
		ad := new(Ad)
		ad.Number = number
		ad.Price = pair.value
		v := ad.IsPriceChange()
		if v != pair.want {
			t.Error(
				"For", pair,
				"expected", pair.want,
				"got", v,
			)
		}
	}
}

