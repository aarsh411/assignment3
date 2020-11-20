package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	BookID      uint   `json:"bookId" gorm:"primary_key"`
	BookName    string `json:"bookName"`
	Description string `json:"description"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "aarsh:12345@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE bookdetail")
	db.Exec("USE bookdetail")
	db.AutoMigrate(&Book{})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/addbook", createBook).Methods("POST")
	router.HandleFunc("/getbook/{bookId}", getBook).Methods("GET")
	router.HandleFunc("/getallbooks", getBooks).Methods("GET")
	router.HandleFunc("/updatebook/{bookId}", updateBook).Methods("PUT")
	router.HandleFunc("/deletebook/{Id}", deleteBook).Methods("DELETE")
	initDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	db.Create(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	db.Preload("Items").Find(&books)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputBookID := params["bookId"]

	var book Book
	db.Preload("Items").First(&book, inputBookID)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	db.Save(&updatedBook)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputBookID := params["bookId"]
	id64, _ := strconv.ParseUint(inputBookID, 10, 64)
	idToDelete := uint(id64)

	db.Where("Book_id = ?", idToDelete).Delete(&Book{})
	w.WriteHeader(http.StatusNoContent)
}
