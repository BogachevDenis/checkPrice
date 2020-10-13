package sender

import (
	"fmt"
	"net/smtp"
)

func SendMail(email []string, newprice, oldprice string)  {
	fmt.Println(email,newprice,oldprice)
	message := []byte("Hello world, new price is " + newprice + ", old price is " + oldprice)
	auth := smtp.PlainAuth(
		"",
		"testavitobuyer-experience@mail.ru",
		"buyer-experience",
		"smtp.mail.ru",
		)
	err := smtp.SendMail(
		"smtp.mail.ru:25",
		auth,
		"testavitobuyer-experience@mail.ru",
		email,
		message)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SubmitEmail(email, key string) error {
	message := []byte("Hello world, you should check this line: http://localhost:8080/"+key)
	auth := smtp.PlainAuth(
		"",
		"testavitobuyer-experience@mail.ru",
		"buyer-experience",
		"smtp.mail.ru",
	)
	err := smtp.SendMail(
		"smtp.mail.ru:25",
		auth,
		"testavitobuyer-experience@mail.ru",
		[]string{email},
		message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}