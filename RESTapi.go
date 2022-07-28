package main

import (
	"Library_Managment_System/Library"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//creating library
var lib = Library.NewLibrary()

//WelcomePage for welcome message
func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Library Management System")
}

// AddMember for adding member
func AddMember(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a user
	var NewMember Library.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	err = json.Unmarshal(reqBody, &NewMember)
	if err != nil {
		fmt.Errorf("can't unmarshal the data")
	}

	//check if member already exists
	Bool := lib.CheckMember(NewMember)
	if Bool == true {
		log.Fatalln("Member Already Exists")
	} else {

		//Add Member to the Library System
		lib.NewMember(NewMember)
	}

	//returns the json object as received to the api
	json.NewEncoder(w).Encode(NewMember)
}

//AddBook for adding member
func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PhysicalBook or DigitalBook[p/d]")

	// get the body of our POST request
	// unmarshal this into a book

	var NewBook map[string]Library.DigitalBook

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	err = json.Unmarshal(reqBody, &NewBook)
	if err != nil {
		fmt.Errorf("can't unmarshal the data")
	}
	for key, value := range NewBook {
		if key == "d" {

			//adding book to the NewDigitalBook
			_, err = Library.NewDigitalBook(value.Name, value.Author, value.Kind, value.Copies)
			if err != nil {
				log.Fatalln("new Book can't be added")

			}

		}
		if key == "p" {
			//adding book to the NewPhysicalBook
			_, err = Library.NewPhysicalBook(value.Name, value.Author, value.Kind)
			if err != nil {
				log.Fatalln("new Book can't be added")

			}

		} else {
			fmt.Errorf("invalid Book type")
		}

		// adding new book to the library
		err = lib.NewBook(&value)
		if err != nil {
			log.Fatalln("error in adding book to library")
		}

		//returns the json object as received to the api
		json.NewEncoder(w).Encode(NewBook)

	}
}

//Borrow for Borrowing book
func Borrow(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a map of user to Book

	var Borrow map[Library.User]string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	err = json.Unmarshal(reqBody, &Borrow)
	if err != nil {
		fmt.Errorf("can't unmarshal the data")
	}
	for key, value := range Borrow {

		//verify the user
		err := lib.CheckMember(Library.User(key))
		if !err {
			log.Fatalln("Not a registered member")

		}

		//verify the book
		book, ok := lib.Checkbook(value)
		if !ok {
			log.Fatalln("Book not available")

		}

		//borrow the book
		error := book.Borrow(key)
		if error != nil {
			log.Fatalln("error in Book issue to the library")

		}
	}
	json.NewEncoder(w).Encode(Borrow)
}

//Return for Returning Book
func Return(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a map of user to book
	var Returns map[Library.User]string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	err = json.Unmarshal(reqBody, &Returns)
	if err != nil {
		fmt.Errorf("can't unmarshal the data")
	}

	//key ,value pair fot the map
	for key, value := range Returns {

		//verify the user
		err := lib.CheckMember(key)
		if !err {
			log.Fatalln("Not a registered member")

		}

		//verify the book
		book, ok := lib.Checkbook(value)
		if !ok {
			log.Fatalln("Book not available")

		}

		//borrow the book
		error := book.Return(key)
		if error != nil {
			log.Fatalln("error in Book issue to the library")

		}
	}
	json.NewEncoder(w).Encode(Returns)
}

//handleRequests contains all the functions necessary for handling API calls using gorilla mux
func handleRequests() {
	//creates a gorilla mux to handle different paths to access variety of functions
	myRouter := mux.NewRouter().StrictSlash(true)
	//homepage to verify the API is working
	myRouter.HandleFunc("/", WelcomePage).Methods("GET")
	//Endpoint to insert member
	myRouter.HandleFunc("/member", AddMember).Methods("POST")
	//Endpoint to insert new book
	myRouter.HandleFunc("/book", AddBook).Methods("POST")
	//Endpoint to borrow a book
	myRouter.HandleFunc("/borrow", Borrow).Methods("POST")
	//Endpoint to return a borrowed book
	myRouter.HandleFunc("/return", Return).Methods("POST")
	//sets the port number to listed to requests
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
