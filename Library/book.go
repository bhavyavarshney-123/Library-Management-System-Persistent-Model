package Library

import "fmt"

// Book interface for different functionalities
type Book interface {
	name() string
	author() string
	kind() BookType
	Borrow(User) error
	Return(User) error
}

// BookType enum to define a book-type
type BookType int

//assigning each BookType a constant value
const (
	eBook BookType = iota
	AudioBook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
	Manga
	SelfHelp
)

//function to check the value and return the designated BookType
func (bookType BookType) String() string {
	switch bookType {
	case eBook:
		return "eBook"
	case AudioBook:
		return "AudioBook"
	case Hardback:
		return "Hardback"
	case Paperback:
		return "Paperback"
	case Encyclopedia:
		return "Encyclopedia"
	case Magazine:
		return "Magazine"
	case Comic:
		return "Comic"
	case Manga:
		return "Manga"
	case SelfHelp:
		return "SelfHelp"
	default:
		return "Unknown"

	}
}

// DigitalBook struct
type DigitalBook struct {
	Name     string   `json:"Name"`
	Author   string   `json:"Author"`
	Copies   int      `json:"Copies"`
	Kind     BookType `json:"Kind"`
	Borrower []User   `json:"Borrower"`
}

func (D *DigitalBook) name() string {
	return D.Name
}
func (D *DigitalBook) author() string {
	return D.Author
}
func (D *DigitalBook) kind() BookType {
	return D.Kind
}
func (D *DigitalBook) Borrow(user User) error {
	//if no copies available to issue
	if len(D.Borrower) >= D.Copies {
		return fmt.Errorf("issue Limit Exceeded")
	}

	//if copies available to issue
	D.Borrower = append(D.Borrower, user)

	return nil
}
func (D *DigitalBook) Return(user User) error {
	//to remove the particular user from the slice of Borrower
	for i, v := range D.Borrower {
		if v == user {
			D.Borrower = append(D.Borrower[:i], D.Borrower[i+1:]...)
			D.Copies++
			i++
			return nil
		}
	}

	//if user not found in the borrower
	return fmt.Errorf("user not have this book")

}

//NewDigitalBook Constructor to add new digital Book
func NewDigitalBook(name, author string, kind BookType, copies int) (*DigitalBook, error) {
	switch kind {
	//only particular kinds of Book can be added
	case eBook, AudioBook, Magazine, Comic, Manga, SelfHelp:
	default:
		return nil, fmt.Errorf("invalid Booktype")
	}

	//Only 200 copies of any Digital Book can be added
	if copies > 200 {
		return nil, fmt.Errorf("can't Store more than 200 Book copies")
	}
	return &DigitalBook{
		name,
		author,
		copies,
		kind, make([]User, 0),
	}, nil
}

// PhysicalBook  struct
type PhysicalBook struct {
	Name     string   `json:"Name"`
	Author   string   `json:"Author"`
	Kind     BookType `json:"Kind"`
	Borrower User     `json:"Borrower"`
}

func (P *PhysicalBook) name() string {
	return P.Name
}
func (P *PhysicalBook) author() string {
	return P.Author
}
func (P *PhysicalBook) kind() BookType {
	return P.Kind
}
func (P *PhysicalBook) Borrow(user User) error {
	//If book already issued
	if P.Borrower != "" {
		return fmt.Errorf("book already issued")
	}
	//new book to issue
	P.Borrower = user
	return nil
}

func (P *PhysicalBook) Return(user User) error {
	//if still Borrower not matched with user then error
	if P.Borrower != user {
		return fmt.Errorf("user not have this book")
	}
	P.Borrower = " "
	return nil
}

//NewPhysicalBook Constructor to add new physical Book
func NewPhysicalBook(name, author string, kind BookType) (*PhysicalBook, error) {
	switch kind {
	//only particular kinds of Book can be added
	case Hardback, Paperback, Encyclopedia, Magazine, Comic, Manga, SelfHelp:
	default:
		return nil, fmt.Errorf("invalid Booktype")
	}
	return &PhysicalBook{
		name,
		author,
		kind, "",
	}, nil
}
