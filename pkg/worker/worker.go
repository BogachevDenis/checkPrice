package worker

import (
	"os"
	"time"
	"github.com/checkPrice/pkg/database"
	"github.com/checkPrice/pkg/sender"
	log "github.com/sirupsen/logrus"
)

func Worker()  {
	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgBase := os.Getenv("PG_BASE")
	err := database.Connect(pgUser,pgPass,pgBase)
	if err != nil {
		log.WithFields(log.Fields{
			"error" : err,
		}).Fatal("Connection to DB error")
	}
	for  {
		ads, err := database.SelectAds()
		if err != nil {
			log.Error(err)
		}
		for _, ad := range ads {
			if ad.IsPriceChange(){
				err := database.UpdatePrice(ad.NewPrice,ad.Id)
				if err != nil{
					log.Error(err)
				}
				emails, err := database.GetEmails(ad.Id)
				if err != nil {
					log.Error(err)
				}
				sender.SendMail(emails,ad.Number,ad.NewPrice,ad.Price)
			}
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}



