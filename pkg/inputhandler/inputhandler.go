package inputhandler

import (
	"math/rand"
	"strings"
	"github.com/asaskevich/govalidator"
	"github.com/checkPrice/pkg/database"
	"github.com/checkPrice/pkg/parser"
	"github.com/checkPrice/pkg/sender"
	log "github.com/sirupsen/logrus"
)

const (
	emptyLine 		string = "Empty line"
	notValidEmail 	string = "It is not a valid email"
	notValidUrl		string = "It is not a valid url"
	emailSubmit		string = "You should submit your email address, please follow the link in the message"
	adNotExist		string = "This ad does not exist"
	addToDbError	string = "Data can`t be sent to DB"
)

type Sub struct{
	Email	string	`json:"email"`
	Url		string	`json:"url"`
	Number 	string 	`json:"number"`
	Price	string 	`json:"price"`
	Error 	string	`json:"error"`
	Message string 	`json:"message"`
}

// Check email
func (s *Sub) CheckEmail() {
	if s.Email == "" {
		s.Error = emptyLine
		return
	}
	s.Email = strings.TrimSpace(s.Email)
	if !govalidator.IsEmail(s.Email){
		s.Error = notValidEmail
		return
	}
}
// Check Url
func (s *Sub) CheckUrl() {
	if s.Error != "" {
		return
	}
	if s.Url == "" {
		s.Error = emptyLine
		return
	}
	s.Url = strings.TrimSpace(s.Url)
	if !govalidator.IsURL(s.Url){
		s.Error = notValidUrl
		return
	}
	if !(strings.HasPrefix(s.Url, "https://www.avito.ru") || strings.HasPrefix(s.Url, "https://avito.ru")|| strings.HasPrefix(s.Url, "https://m.avito.ru")){
		s.Error = notValidUrl
		return
	}
}
// Add number
func (s *Sub) CreateNumber() {
	if s.Error != "" {
		return
	}
	length := len(s.Url) - 10
	s.Number = s.Url[length:]
}
// Add price
func (s *Sub) AddPrice() {
	if s.Error != "" {
		return
	}
	s.Price, _ = parser.GetPrice(s.Number)
	if s.Price == "" {
		s.Error = adNotExist
	}
}

// Add data to database
func (s *Sub) AddSubscription() {
	if s.Error != "" {
		return
	}
	key := RandStringRunes(15)
	if !database.IsEmailExists(s.Email) {
		s.Message = emailSubmit
		sender.SubmitEmail(s.Email,key)
	}
	err := database.InsertData(s.Email, s.Number, s.Price, key)
	if err != nil{
		s.Error = addToDbError
		log.WithFields(log.Fields{
			"data" : s,
			"error": err,
		}).Error("Data can`t be sent to DB")
		return
	}
	log.WithFields(log.Fields{
		"data" : s,
	}).Info("Subscription was created")
}

// Generate random key
func RandStringRunes(n int) string {
	var Runes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXZYabcdefghijklmnopqrstuvwxyz0123456789")
	mas := make([]rune, n)
	for i := range mas {
		mas[i] = Runes[rand.Intn(len(Runes))]
	}
	return string(mas)
}

