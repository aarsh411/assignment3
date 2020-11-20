package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add/{a}/{b}", Add).Methods("GET")
	router.HandleFunc("/sub/{a}/{b}", Sub).Methods("GET")
	router.HandleFunc("/mul/{a}/{b}", Mul).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", Div).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
func Add(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a := params["a"]
	int1, err := strconv.ParseInt(a, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	b := params["b"]

	int2, err := strconv.ParseInt(b, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	var c int64 = int1 + int2
	//int3, err := strconv.ParseInt(c, 6, 12)
	fmt.Printf("result is %d \n", c)
}

func Sub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a := params["a"]
	int1, err := strconv.ParseInt(a, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	b := params["b"]

	int2, err := strconv.ParseInt(b, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	var c int64 = int1 - int2
	//int3, err := strconv.ParseInt(c, 6, 12)
	fmt.Printf("result is %d \n", c)
}
func Mul(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a := params["a"]
	int1, err := strconv.ParseInt(a, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	b := params["b"]

	int2, err := strconv.ParseInt(b, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	var c int64 = int1 * int2
	//int3, err := strconv.ParseInt(c, 6, 12)
	fmt.Printf("result is %d \n", c)
}

func Div(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a := params["a"]
	int1, err := strconv.ParseInt(a, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	b := params["b"]

	int2, err := strconv.ParseInt(b, 6, 12)
	if err != nil {
		fmt.Println(err)
		//panic("failed to connect database")
	}
	var c int64 = int1 / int2
	//int3, err := strconv.ParseInt(c, 6, 12)
	fmt.Printf("result is %d \n", c)
}
