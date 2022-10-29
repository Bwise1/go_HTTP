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
	godotenv.Load()
	port :=os.Getenv("PORT")


	r.HandleFunc("/", task).Methods("GET")
	address := fmt.Sprintf(":%s",port)

	fmt.Println("Server started at port ",address)
	log.Fatal(http.ListenAndServe(address, r))
}
