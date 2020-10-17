package main

import (
	"encoding/json"
	"github.com/checkPrice/pkg/inputhandler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateSubscription(t *testing.T)  {
	var sub inputhandler.Sub
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ `{"email":"","url":"test"}`,				"Empty line" },
		{ `{"email":"test@mail","url":"test"}`,		"It is not a valid email" },
		{ `{"email":"test@mail.ru","url":""}`,		"Empty line" },
		{ `{"email":"test@mail.ru","url":"test"}`,	"It is not a valid url" },
		{ `{"email":"test@mail.ru","url":"https://wito.ru/stupino/chasy_i_ukrasheniya/chasy_casio_g-shock_ga-100b_2003091497"}`,	"It is not a valid url" },
		{ `{"email":"dsgd@mail.ru","url":"https://www.avito.ru/stupino/chasy_i_ukrasheniya/chasy_casio_g-shock_ga-100b_20030"}`,	"This ad does not exist" },
	}
	for _, pair := range tests {
		r := strings.NewReader(pair.value)
		req, err := http.NewRequest("POST", "/create", r)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(createSubscription)
		handler.ServeHTTP(rr, req)
		body, _ := ioutil.ReadAll(rr.Body)
		json.Unmarshal(body, &sub)
		if sub.Error != pair.want {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), pair.want)
		}

	}
}