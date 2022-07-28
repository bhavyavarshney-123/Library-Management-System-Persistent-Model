package main

import (
	"Library_Managment_System/Library"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var input int
	flag := true
	lib := Library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)
	//flag to Choose the required action to perform
	for flag {
		fmt.Println("Welcome to Library Management System")
		fmt.Println("1. Register New Member")
		fmt.Println("2. Add Book to the Library")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Enter API")
		fmt.Println("6. Exit")
		fmt.Printf("Enter your choice: ")
		if _, err := fmt.Scanln(&input); err != nil {
			log.Fatalln("scan failed", err)
		}

		switch input {
		//Register a User
		case 1:

			fmt.Println("Enter Your Name:")
			//scan the name
			scanner.Scan()
			name := scanner.Text()

			//check if already a user
			err := lib.CheckMember(Library.User(name))
			if !err {
				log.Fatalln("Already a member")

			}

			//add new member to the Library
			lib.NewMember(Library.User(name))
			fmt.Printf("Welcome %s \n\n", name)

		//Add Book
		case 2:

			//enter book name
			fmt.Println("Enter Book Name:")
			scanner.Scan()
			bookname := scanner.Text()

			//verify if book already registered
			book, ok := lib.Checkbook(bookname)
			if !ok {
				log.Fatalln("Book already registered")

			}

			//Enter author name
			fmt.Println("Enter Author Name:")
			scanner.Scan()
			author := scanner.Text()

			//Enter booktype
			fmt.Println("Enter Book Type:")
			var kind Library.BookType
			if _, err := fmt.Scanln(&kind); err != nil {
				log.Fatalln("scan failed", err)
			}

			//check for book digital or physical
			fmt.Println("PhysicalBook or DigitalBook[p/d]")
			scanner.Scan()
			check := scanner.Text()

			var err error
			//if it's a digital book

			if check == "d" {
				//input number of copies to add
				var Copies int
				fmt.Println("Enter number of Copies:")
				if _, err := fmt.Scanln(&Copies); err != nil {
					log.Fatalln("scan failed", err)

				}

				//adding book to the NewDigitalBook
				book, err = Library.NewDigitalBook(bookname, author, kind, Copies)
				if err != nil {
					log.Fatalln("new Book can't be added")

				}

			} else {
				//adding book to the NewPhysicalBook
				book, err = Library.NewPhysicalBook(bookname, author, kind)
				if err != nil {
					log.Fatalln("new Book can't be added")

				}

			}

			// adding new book to the library
			err = lib.NewBook(book)
			if err != nil {
				log.Fatalln("error in adding book to library")
			}

			fmt.Print("Book added Successfully \n\n")

		//Borrow Book
		case 3:
			//enter member
			fmt.Println("Enter your Name:")
			scanner.Scan()
			name := scanner.Text()

			//verify the user
			err := lib.CheckMember(Library.User(name))
			if !err {
				log.Fatalln("Not a registered member")

			}

			//enter book name
			fmt.Println("Enter your Book Name:")
			scanner.Scan()
			BookName := scanner.Text()

			//verify the book
			book, ok := lib.Checkbook(BookName)
			if !ok {
				log.Fatalln("Book not available")

			}

			//borrow the book
			error := book.Borrow(Library.User(name))
			if error != nil {
				log.Fatalln("error in Book issue to the library")

			}

			fmt.Print("Book issued Successfully \n\n")

			//return Book
		case 4:
			//enter member
			fmt.Println("Enter your Name:")
			scanner.Scan()
			name := scanner.Text()

			//verify the user
			err := lib.CheckMember(Library.User(name))
			if !err {
				log.Fatalln("Not a registered member")

			}

			//enter book name
			fmt.Println("Enter your Book Name:")
			scanner.Scan()
			BookName := scanner.Text()

			//verify the book
			book, ok := lib.Checkbook(BookName)
			if !ok {
				log.Fatalln("Book not available")

			}

			//return book
			error := book.Return(Library.User(name))
			if error != nil {
				log.Fatalln("error in returning book")

			}

			fmt.Print("Return Successful \n\n")

			//Sets the program to begin accepting api calls on localhost:10000
		case 5:
			handleRequests()
			//exit clause to close application

		default:
			flag = false
		}

	}
}
