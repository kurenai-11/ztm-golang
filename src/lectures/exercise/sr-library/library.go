//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name         string
	lended       bool
	lendedBy     *Member
	checkinTime  int64
	checkoutTime int64
}

type Member struct {
	name        string
	lendedBooks []*Book
}

type Library struct {
	books   []Book
	members []Member
}

var library Library

func checkinBook(book *Book, member *Member) {
	book.lendedBy = member
	book.lended = true
	book.checkinTime = time.Now().Unix()
	member.lendedBooks = append(member.lendedBooks, book)
}

func checkoutBook(book *Book) {
	previousLender := *book.lendedBy
	book.lendedBy = nil
	book.lended = false
	book.checkoutTime = time.Now().Unix()
	newBooks := make([]*Book, 0)
	for _, lendedBook := range previousLender.lendedBooks {
		if lendedBook.name != book.name {
			newBooks = append(newBooks, lendedBook)
		}
	}
	previousLender.lendedBooks = newBooks
}

func main() {
	library.books = append(library.books, Book{name: "Harry Potter"}, Book{name: "Sherlock Holmes"}, Book{name: "Batman"}, Book{name: "Rave"})
	library.members = append(library.members, Member{name: "Harry"}, Member{name: "Sherlock"}, Member{name: "Bruce"})
	checkinBook(&library.books[0], &library.members[0])
	checkinBook(&library.books[1], &library.members[1])
	checkinBook(&library.books[2], &library.members[2])
	checkinBook(&library.books[3], &library.members[2])
	fmt.Println(library)
	checkoutBook(&library.books[0])
	fmt.Println(library)
}
