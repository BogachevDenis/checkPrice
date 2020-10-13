package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)
var db *sql.DB

type Ad struct{
	Id		int
	Number		string
	Price	string
}

func Connect (pg_host, pg_port, pg_user, pg_pass, pg_base string ) error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",pg_host, pg_port, pg_user, pg_pass, pg_base))
	if err != nil {
		fmt.Println("errOpen",err.Error())
		return err
	}
	return nil
}

func IsEmailExists(email string) bool {
	r := db.QueryRow(`SELECT "id" FROM "users" WHERE "email" = $1`, email)
	var id int
	err := r.Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func InsertData(email, number, price, key string) error {
	userId, err := getEmailId(email, key)
	if err != nil{
		return err
	}
	fmt.Println("emailid =", userId)
	numberId, err := getAdId(number, price)
	if err != nil{
		return err
	}
	fmt.Println("numberid =", numberId)
	subId, err := getSubId(userId, numberId)
	if err != nil{
		return err
	}
	fmt.Println("subId =", subId)
	return nil
}



func getEmailId(email, key string) (int, error) {
	r := db.QueryRow(`SELECT "id" FROM "users" WHERE "email" = $1`, email)
	var id int
	err := r.Scan(&id)
	if err == sql.ErrNoRows {
		k := db.QueryRow(`INSERT INTO users (email, key) VALUES ($1, $2) returning id`, email, key)
		err := k.Scan(&id)
		if err != nil {
			log.WithFields(log.Fields{
				"error" : err,
			}).Error("Insert email error")
		}
	}
	return id, nil
}

func getAdId(number string, price string) (int,error) {
	r := db.QueryRow(`SELECT "id" FROM "ads" WHERE "number" = $1`, number)
	var id int
	err := r.Scan(&id)
	if err == sql.ErrNoRows {
		k := db.QueryRow(`INSERT INTO ads (number, price) VALUES ($1, $2) returning id`, number, price)
		err := k.Scan(&id)
		if err != nil {
			log.WithFields(log.Fields{
				"error" : err,
			}).Error("Insert url error")
		}
	}
	return id, nil
}

func getSubId(userId, numberId int) (int, error) {
	r := db.QueryRow(`SELECT "id" FROM "subscription" WHERE "userid" = $1 AND "adid" = $2`, userId, numberId)
	var id int
	err := r.Scan(&id)
	if err == sql.ErrNoRows {
		k := db.QueryRow(`INSERT INTO subscription (userid, adid) VALUES ($1, $2) returning id`, userId, numberId)
		err := k.Scan(&id)
		if err != nil {
			log.WithFields(log.Fields{
				"error" : err,
			}).Error("Insert subscription error")
		}
	}
	return id, nil
}





func SelectAds() ([]*Ad, error) {
	rows, err := db.Query(`SELECT * FROM "ads"`)
	if err != nil {
		return []*Ad{}, err
	}
	defer rows.Close()
	ads := make([]*Ad, 0)
	for rows.Next() {
		ad := new(Ad)
		err := rows.Scan(&ad.Id,&ad.Number,&ad.Price)
		if err != nil {
			return []*Ad{}, err
		}
		ads = append(ads, ad)
	}
	return ads, nil
}

func UpdatePrice(price string, id int) error {
	_, err := db.Exec(`UPDATE "ads" SET price = $1 WHERE id = $2;`,price, id )
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetEmails(adid int) ([]string,error) {
	rows, err := db.Query(`SELECT u.email FROM "users" u INNER JOIN "subscription" s ON u.id = s.userid AND u.activation = true AND s.adid = $1`, adid)
	if err != nil {
		fmt.Println("err 2 =", err)
		return []string{""}, nil
	}
	defer rows.Close()
	emails := make([]string, 0)
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			fmt.Println("err 3 =", err)
			return []string{""}, nil
		}
		emails = append(emails, email)
	}
	return emails, nil
}

func UpdateActivation(key string) error {
	_, err := db.Exec(`UPDATE "users" SET activation = true WHERE key = $1;`,key )
	if err != nil {
		fmt.Println(err)
	}
	return nil
}