package main

import (
	"./lib"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./webroot")))
	http.HandleFunc("/api/topics", topics)
	http.HandleFunc("/topic/", serveIndex)
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error())
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webroot/index.html")
}

func topics(w http.ResponseWriter, r *http.Request) {
	controller := lib.Controller{}
	switch r.Method {
	case "POST":
		topic := lib.Topic {}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &topic)
		controller.InsertTopic(topic)
	case "PUT":
		topic := lib.Topic {}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &topic)
		controller.UpdateTopic(topic)
	case "GET":
		w.Write(controller.GetTopics())
	}
}
