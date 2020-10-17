package parser

import (
	"testing"

)

func TestIParser(t *testing.T)  {

	type testpair struct {
		value 	string
		want 	string
	}

	var tests = []testpair{
		{ "1862432554",	"4000" },
	}
	for _, pair := range tests {
		v, _ :=GetPrice(pair.value)

		if v != pair.want {
			t.Error(
				"For", pair,
				"expected", pair.want,
				"got", v,
			)
		}
	}
}
