package Library

import (
	"Library_Managment_System/badgerdb"
	"fmt"
	"github.com/dgraph-io/badger"
)

type User string

//Library struct for members and BookDetails
type Library struct {
	members     map[User]struct{}
	BookDetails map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		make(map[User]struct{}),
		make(map[string]Book),
	}
}

//NewMember to add new member to the library
func (L *Library) NewMember(user User) {
	L.members[user] = struct{}{}

	////adding member to the database
	err := L.MemberDB(user)
	if err != nil {
		fmt.Errorf("member not added to Database")
	}

}

//NewBook to add new book to the library
func (L *Library) NewBook(book Book) error {
	if _, exists := L.BookDetails[book.name()]; exists {
		return fmt.Errorf("'%v' Book already exists ", book.name())
	}
	L.BookDetails[book.name()] = book

	////adding book to the database
	err := L.BookDb(book)
	if err != nil {
		fmt.Errorf("book not added to Database")
	}

	return nil
}

//CheckMember to check if a particular user exits or not
func (L *Library) CheckMember(user User) bool {
	//Check member in the cache memory
	_, exists := L.members[user]

	//if present then return true
	if exists == true {
		return exists
	} else {
		fmt.Print("checking in db \n\n")
		//else check in the database
		data := []byte(user)
		err, exists := L.GetMemberDb(data)
		if err != nil {
			fmt.Errorf("error in getting member from the database ")
		}
		return exists
	}

}

//Checkbook to check about a particular book
func (L *Library) Checkbook(BookName string) (Book, bool) {
	//Check Book in the cache memory
	book, ok := L.BookDetails[BookName]

	//if present then return true
	if ok == true {
		return book, ok
	} else {
		//else check in the database
		data := []byte(BookName)
		book, ok := L.GetBookDb(data)
		return book, ok
	}
}

// MemberDB to store the member in the database
func (*Library) MemberDB(user User) error {
	//Opening the Badger database
	db, err := badgerdb.Open()
	if err != nil {
		return fmt.Errorf("error Opening the Database")
	}

	//converting the user into slice of bytes
	data := []byte(user)

	//Storing the data the in the db as user as a key and empty slice of bytes as a value
	err = db.SetEntry(data, []byte{})
	if err != nil {
		return fmt.Errorf("can't store in database")
	}

	//closing the db
	defer db.Close()
	return nil
}

//BookDb to store the book in the database
func (*Library) BookDb(book Book) error {
	db, err := badgerdb.Open()
	if err != nil {
		return fmt.Errorf("error Opening the Database")
	}

	//converting the book name as key to store in the db
	data := []byte(book.name())

	//Serializing the Book
	SerializedBook, err := GobEncode(book)
	if err != nil {
		return fmt.Errorf("can't Serialize the data")
	}

	//Storing the data the in the db as Book name as a key and book as a value
	er := db.SetEntry(data, SerializedBook)
	if er != nil {
		return fmt.Errorf("can't store in database")
	}

	//closing the db
	defer db.Close()
	return nil
}

//GetMemberDb gives the member information as a key
func (*Library) GetMemberDb(key []byte) (error, bool) {
	//Opening the Badger database
	db, err := badgerdb.Open()
	if err != nil {
		return fmt.Errorf("error Opening the Database"), false
	}

	//checking for a particular key by defining a view method
	var check bool
	err = db.Client.View(func(view *badger.Txn) error {
		//Attempt to get the Item for the given key
		_, err := view.Get(key)
		if err != nil {
			check = false
			return fmt.Errorf("db get on Key '%x' fail:'%w'", key, err)
		}
		check = true
		return nil
	})

	//Closing the db
	defer db.Close()
	return err, check
}

//GetBookDb gives the details of the book as value when provided with book name as key
func (*Library) GetBookDb(key []byte) (Book, bool) {
	db, err := badgerdb.Open()
	if err != nil {
		fmt.Errorf("error Opening the Database")
		return nil, false
	}

	//if key not found in the database then returning the error
	value, err := db.GetEntry(key)
	if err != nil {
		fmt.Errorf("error in finding key in the Db")
		return nil, false

	}

	//Deserialized the data into the book object
	var object Book
	book, err := GobDecode(value, object)
	if err != nil {
		fmt.Errorf("book can't be Deserialized")
		return nil, false
	}

	//close the db
	defer db.Close()
	return book, true
}
