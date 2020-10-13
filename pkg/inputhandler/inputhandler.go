package inputhandler

import (
	"fmt"
	"github.com/checkPrice/pkg/parser"
	"github.com/checkPrice/pkg/database"
	"github.com/asaskevich/govalidator"
	"github.com/checkPrice/pkg/sender"
	"math/rand"
	"strings"
)

const (
	emptyLine 		string = "Empty line"
	notValidEmail 	string = "It is not a valid email"
	notValidUrl		string = "It is not a valid url"
	emailSubmit		string = "You should submit your email address, please follow the link in the message"
)

type Sub struct{
	Email	string	`json:"email"`
	Url		string	`json:"url"`
	Number 	string 	`json:"number"`
	Price	string 	`json:"price"`
	Error 	string	`json:"error"`
}


func (s *Sub) CheckEmail() {
	if s.Email == "" {
		s.Error = emptyLine
		return
	}
	if !govalidator.IsEmail(s.Email){
		s.Error = notValidEmail
		return
	}
}

func (s *Sub) CheckUrl() {
	if s.Error != "" {
		return
	}
	if s.Url == "" {
		s.Error = emptyLine
		return
	}
	if !govalidator.IsURL(s.Url){
		s.Error = notValidUrl
		return
	}
	if !(strings.HasPrefix(s.Url, "https://www.avito.ru") || strings.HasPrefix(s.Url, "https://avito.ru")|| strings.HasPrefix(s.Url, "https://m.avito.ru")){
		s.Error = notValidUrl
		return
	}
}

func (s *Sub) CreateNumber() {
	if s.Error != "" {
		return
	}
	length := len(s.Url) - 10
	s.Number = s.Url[length:]
}

func (s *Sub) AddPrice() {
	if s.Error != "" {
		return
	}
	s.Price = parser.GetPrice(s.Number)
}



func (s *Sub) AddSubscription() {
	if s.Error != "" {
		return
	}
	key := RandStringRunes(11)
	if !database.IsEmailExists(s.Email) {
		s.Error = emailSubmit
		sender.SubmitEmail(s.Email,key)
	}
	err := database.InsertData(s.Email, s.Number, s.Price, key)
	if err != nil{
		fmt.Println(err)
		return
	}
}

// Generate random key
func RandStringRunes(n int) string {
	var Runes = []rune("ABCDEFGHIWXZYabcdefghijklmnopqrstuvwxyz0123456789")
	mas := make([]rune, n)
	for i := range mas {
		mas[i] = Runes[rand.Intn(len(Runes))]
	}
	return string(mas)
}

