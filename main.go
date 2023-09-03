package main

import (
	"encoding/json" //to read json file
	"fmt"
	"log" //log error

	"math/rand" //to generate random nuber
	"net/http"

	"strconv" //string converter

	"github.com/gorilla/mux" //router
)

// struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
}

// init books var as slice
var books []Book

//get all books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the parameter

	//loop through books and find id
	for _, item := range books { // '_' is the index
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) //generate random number between (1-100)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// update content of the book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // access to parameter
	//loop through the bokks
	// Delete the movie that nedded to be updated
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			//add a new book - that needed to be updated
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // use the same id given by the user
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

// delte book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // acces to prameter
	//loop through the movies
	for index, item := range books {
		if item.ID == params["id"] {
			// the index user given the index is replaced in the slice
			// and remaing indexes are adedd to the slice again
			// books[:index] -> index that nedded to be deleted
			// books[index+1:...] -. represents the remaning bokks that are there in the books slice
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

func main() {

	//init router

	router := mux.NewRouter()

	//data
	books = append(books, Book{ID: "1", ISBN: "86348", Title: "Book1", Author: &Author{FName: "Ryan", LName: "Gosling"}})
	books = append(books, Book{ID: "2", ISBN: "85578", Title: "Book2", Author: &Author{FName: "Arup", LName: "Das"}})

	//route handelers /endpoits handeer

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//to run server
	// fmt.Printf("Starting server at port 8000\n")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Starting port 9000")
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal(err)
	}
	// as ListenAndServe returns non nil error so
	// line 93 can also be written as
	//log.Fatal(http.ListenAndServe(":9000", router))
}
