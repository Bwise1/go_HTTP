package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	data := &details{SlackUsername:"MayorBenjys", Backend:true, Age: 27, Bio: "I am new to this part of the tech world, hoping to make it through to the end of the internship"}
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
	godotenv.Load()
	port :=os.Getenv("PORT")


	r.HandleFunc("/", task).Methods("GET")
	address := fmt.Sprintf(":%s",port)

	fmt.Println("Server started at port ",address)
	log.Fatal(http.ListenAndServe(address, r))
}
