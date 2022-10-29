package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type details struct {
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int `json:"age"`
	Bio string `json:"bio"`
}

func task(w http.ResponseWriter, req *http.Request) {
	data := &details{SlackUsername:"Benjys", Backend:true, Age: 27, Bio: "I don't want peace, I am a fight ..."}
	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Write(jsonData)


}


func main() {

	r := mux.NewRouter()



	r.HandleFunc("/", task).Methods("GET")

	fmt.Println("Server started at port 3001")
	log.Fatal(http.ListenAndServe(":8090", r))
}
