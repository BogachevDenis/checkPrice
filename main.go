package main

import (
	"encoding/json"
	"fmt"
	"github.com/checkPrice/pkg/database"
	"github.com/checkPrice/pkg/inputhandler"
	"github.com/checkPrice/pkg/worker"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func init()  {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	go worker.Worker()
}

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/",mainPage)
	r.HandleFunc("/{key}", submition).Methods("GET")
	r.HandleFunc("/create", createSubscription).Methods("POST")

	log.WithFields(log.Fields{
		"port" : port,
	}).Info("Starting Server")

	http.ListenAndServe(":" + port,r)
}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func createSubscription(w http.ResponseWriter, r *http.Request)  {
	sub := new(inputhandler.Sub)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"file" : r.Body,
			"error": err,
		}).Warning("Read request error")
	}
	err = json.Unmarshal(body, &sub)
	if err != nil {
		log.WithFields(log.Fields{
			"file" : body,
			"error": err,
		}).Warning("Unmarshal error")
	}
	sub.CheckEmail()
	sub.CheckUrl()
	sub.CreateNumber()
	sub.AddPrice()
	sub.AddSubscription()
	savedSub , err := json.Marshal(sub)
	if err == nil {
		log.WithFields(log.Fields{
			"data" : sub,
		}).Info("Creating subscription")
	}
	w.WriteHeader(200)
	w.Write(savedSub)
}

func submition(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	key := vars["key"]
	err := database.UpdateActivation(key)
	if err == nil{
		w.Write([]byte("your mail is confirmed, subscription created "))
	}
}