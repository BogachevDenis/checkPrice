package sender

import (
	"net/smtp"
	"os"
	log "github.com/sirupsen/logrus"
)

const (
	senderEmail 	string = "testavitobuyer-experience@mail.ru"
	senderPassword 	string = "buyer-experience"
	host			string = "smtp.mail.ru"
	addr 			string = "smtp.mail.ru:25"
)

func SendMail(email []string, number, newprice, oldprice string) {
	message := []byte("Hello, your ad - "+ number +" changed price, new price is " + newprice + ", old price is " + oldprice)
	auth := smtp.PlainAuth(
		"",
		senderEmail,
		senderPassword,
		host,
		)
	err := smtp.SendMail(
		addr,
		auth,
		senderEmail,
		email,
		message)
	if err != nil {
		log.Error(err)
		return
	}
}

func SubmitEmail(email, key string) {
	confirmationUrl := os.Getenv("CONFIRMURL")
	message := []byte("To confirm email you should run this line:" + confirmationUrl + key)
	auth := smtp.PlainAuth(
		"",
		senderEmail,
		senderPassword,
		host,
	)
	err := smtp.SendMail(
		addr,
		auth,
		senderEmail,
		[]string{email},
		message)
	if err != nil {
		log.Error(err)
	}
	return
}