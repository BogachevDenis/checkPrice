package database

import (
	"database/sql"
	"fmt"
	"github.com/checkPrice/pkg/adhandler"
	log "github.com/sirupsen/logrus"
)


var db *sql.DB

func Connect (pg_user, pg_pass, pg_base string ) error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", pg_user, pg_pass, pg_base))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func IsEmailExists(email string) bool {
	r := db.QueryRow(`SELECT "id" FROM "users" WHERE "email" = $1 AND activation = true `, email)
	var id int
	err := r.Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func InsertData(email, number, price, key string) error {
	err := insertEmail(email, key)
	if err != nil{
		return err
	}
	err = insertAd(number, price)
	if err != nil{
		return err
	}

	err = insertSub(email, number)
	if err != nil{
		return err
	}
	return nil
}



func insertEmail(email, key string) error {
	_, err := db.Exec(`INSERT INTO users (email, key) VALUES ($1, $2)`, email, key)
	if err != nil {
		_, err = db.Exec(`UPDATE "users" SET key = $1 WHERE email = $2`, key, email)
		if err != nil {
			log.Error("Update key error:", err)
		}
	}
	return nil
}

func insertAd(number string, price string) error {
	_, err := db.Exec(`INSERT INTO ads (number, price) VALUES ($1, $2)`, number, price)
	if err != nil {
		log.Error("Insert ad error")
	}
	return nil
}

func insertSub(email, number string) error {
	_, err := db.Exec(`INSERT INTO subscription (userid, adid) VALUES ((SELECT "id" FROM "users" WHERE "email" = $1),(SELECT "id" FROM "ads" WHERE "number" = $2)) returning id`, email, number)
	if err != nil {
		log.Error("Insert subscription error")
	}
	return nil
}





func SelectAds() ([]*adhandler.Ad, error) {
	rows, err := db.Query(`SELECT * FROM "ads"`)
	if err != nil {
		return []*adhandler.Ad{}, err
	}
	defer rows.Close()
	ads := make([]*adhandler.Ad, 0)
	for rows.Next() {
		ad := new(adhandler.Ad)
		err := rows.Scan(&ad.Id,&ad.Number,&ad.Price)
		if err != nil {
			return []*adhandler.Ad{}, err
		}
		ads = append(ads, ad)
	}
	return ads, nil
}

func UpdatePrice(price string, id int) error {
	_, err := db.Exec(`UPDATE "ads" SET price = $1 WHERE id = $2;`,price, id )
	if err != nil {
		log.Error("UpdatePrice error", err)
	}
	return nil
}

func GetEmails(adid int) ([]string,error) {
	rows, err := db.Query(`SELECT u.email FROM "users" u INNER JOIN "subscription" s ON u.id = s.userid AND u.activation = true AND s.adid = $1`, adid)
	if err != nil {
		log.Error("GetEmails Query error:", err)
		return []string{""}, nil
	}
	defer rows.Close()
	emails := make([]string, 0)
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Error("GetEmails Scan rows error:", err)
			return []string{""}, nil
		}
		emails = append(emails, email)
	}
	return emails, nil
}

func UpdateActivation(key string) error {
	_, err := db.Exec(`UPDATE "users" SET activation = true WHERE key = $1;`,key )
	if err != nil {
		log.Error(err)
	}
	return nil
}
