package worker

import (
	"fmt"
	"github.com/checkPrice/pkg/database"
	"github.com/checkPrice/pkg/parser"
	"github.com/checkPrice/pkg/sender"
	"time"
)

func Worker()  {
	database.Connect("localhost","5432","postgres","1234","subsribe")
	for  {
		ads, _ := database.SelectAds()
		for i, _ := range ads {
			newPrice := parser.GetPrice(ads[i].Number)
			if (ads[i].Price != newPrice) && (newPrice != "") {
				err := database.UpdatePrice(newPrice,ads[i].Id)
				if err != nil{
					fmt.Println("err1 = ",err)
				}

				emails, _ := database.GetEmails(ads[i].Id)
				sender.SendMail(emails,newPrice,ads[i].Price)
			}
			fmt.Println(ads[i])
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}