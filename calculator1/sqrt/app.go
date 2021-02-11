package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type Operands struct {
	OperandOne float32 `json:"operandOne,string"`
	OperandTwo float32 `json:"operandTwo,string"`
}

func sqrt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	json.NewDecoder(r.Body).Decode(&operands)
	fmt.Println(fmt.Sprintf("%s%f", "Getting Sqrt ", operands.OperandOne))
	json.NewEncoder(w).Encode(math.Sqrt(float64(operands.OperandOne)))
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sqrt", sqrt).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6000", router))
}
