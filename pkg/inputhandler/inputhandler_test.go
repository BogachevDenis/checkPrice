package inputhandler

import (
	"github.com/checkPrice/pkg/database"
	"os"
	"testing"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func TestCheckEmail(t *testing.T) {
	sub := new(Sub)
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ "",				"Empty line" },
		{ "test@",		"It is not a valid email" },
		{ "test@mail.",	"It is not a valid email" },
		{ "test@mail.ru",	"" },
	}
	for _, pair := range tests {
		sub.Error = ""
		sub.Email = pair.value
		sub.CheckEmail()
		v := sub.Error
		if v != pair.want {
			t.Error(
				"For", pair,
				"expected", pair.want,
				"got", v,
			)
		}
	}
}

func TestCheckUrl(t *testing.T) {
	sub := new(Sub)
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ "",																		"Empty line" },
		{ "www.avito.ru/kolomna/kollektsionirovanie/zerkalo_1892220447",			"It is not a valid url" },
		{ "wwwavitoru/kolomna/kollektsionirovanie/zerkalo_1892220447",			"It is not a valid url" },
		{ "https://w.avito.ru/kolomna/kollektsionirovanie/zerkalo_1892220447",	"It is not a valid url" },
		{ "https://www.avito.ru/kolomna/kollektsionirovanie/zerkalo_1892220447",	"" },
	}
	for _, pair := range tests {
		sub.Error = ""
		sub.Url = pair.value
		sub.CheckUrl()
		v := sub.Error
		if v != pair.want {
			t.Error(
				"For", pair,
				"expected", pair.want,
				"got", v,
			)
		}
	}
}


func TestCreateNumber(t *testing.T) {
	sub := new(Sub)
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ "www.avito.ru/kolomna/kollektsionirovanie/zerkalo_1892220447",	"1892220447" },
		{ "wwwavitoru/kolomna/kollektsionirovanie/zerkalo_1892220447",	"1892220447" },
	}
	for _, pair := range tests {
		sub.Url = pair.value
		sub.CreateNumber()
		v := sub.Number
		if v != pair.want {
			t.Error(
				"For", pair,
				"expected", pair.want,
				"got", v,
			)
		}
	}
}

func TestAddPrice(t *testing.T) {
	sub := new(Sub)
	type testpair struct {
		value 	string
		notwant 	string
	}
	var tests = []testpair{
		{ "https://www.avito.ru/stupino/mebel_i_interer/stol_2040858409","" },

	}
	for _, pair := range tests {
		sub.Url = pair.value
		sub.CreateNumber()
		sub.AddPrice()
		v := sub.Price
		if v == pair.notwant {
			t.Error(
				"For", pair,
				"not expected", pair.notwant,
				"got", v,
			)
		}
	}
}

func TestAddSubscription(t *testing.T) {
	godotenv.Load("../../.env")
	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgBase := os.Getenv("PG_BASE")
	database.Connect(pgUser,pgPass,pgBase)
	sub := new(Sub)
	type testpair struct {
		emailValue 	string
		urlValue 	string
		want 		string
	}
	var tests = []testpair{
		{ "testavitobuyer-experience@mail.ru","https://www.avito.ru/stupino/mebel_i_interer/stol_2040858409","" },

	}
	for _, pair := range tests {
		sub.Error = ""
		sub.Email = pair.emailValue
		sub.Url = pair.urlValue
		sub.CheckEmail()
		sub.CheckUrl()
		sub.CreateNumber()
		sub.AddPrice()
		sub.AddSubscription()
		v := sub.Error
		if v != pair.want {
			t.Error(
				"For", pair,
				"not expected", pair.want,
				"got", v,
			)
		}
	}
}


