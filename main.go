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

type Operator string

const(
	addition = iota+1
	subraction
	multiplication
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

type response struct{
	SlackUsername string `json:"slackUsername"`
	Operation_type Operator `json:"operation_type"`
	Result int `json:"result"`
}

type Equation struct {
	X int `json:"x"`
	Y int `json:"y"`
	Operation_type Operator `json:"operation_type"`
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

func task2(w http.ResponseWriter, req *http.Request){
	equ := Equation{}
	//var equation Equation
    //json.Unmarshal(reqBody)
	json.NewDecoder(req.Body).Decode(&equ);

	if(equ.Operation_type =="multiplication" || equ.Operation_type =="addition" ||  equ.Operation_type =="subraction" ){
		answer:= calc(equ.X,equ.Y,equ.Operation_type)
		fmt.Println(answer)
		resp:= &response{SlackUsername: "MayorBenjys", Operation_type: equ.Operation_type, Result: answer}
		enableCors(&w)

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}
		w.Write(jsonData)
	}


}
func calc(number1 int, number2 int, operator Operator) int {
	output := 0
	switch operator {
        case "addition":
            output = number1 + number2
        case "subtraction":
            output = number1 - number2
        case "multiplication":
            output = number1 * number2
        default:
            fmt.Println("Invalid Operation")
    }
	return output
}


func main() {

	r := mux.NewRouter()
	godotenv.Load()
	port :=os.Getenv("PORT")


	r.HandleFunc("/", task).Methods("GET")
	r.HandleFunc("/", task2).Methods("POST")

	address := fmt.Sprintf(":%s",port)

	fmt.Println("Server started at port ",address)
	log.Fatal(http.ListenAndServe(address, r))
	//fmt.Println(calc(2,4, addition))
}
