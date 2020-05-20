package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func updateBook(book *Books, author string) {
	book.author = author
}

func printBook(book *Books) {
	fmt.Println("title:", book.title)
	fmt.Println("author:", book.author)
	fmt.Println("subject:", book.subject)
	fmt.Println("book_id:", book.book_id)
}

func main() {

	fmt.Println("hello")

	var book1 Books
	var book2 Books

	book1.title = "go go go"
	book1.author = "www.runoob.com"
	book1.subject = "Go book"
	book1.book_id = 64782

	book2.title = "python book"
	book2.author = "fan pf"
	book2.subject = "python language"
	book2.book_id = 1111222

	printBook(&book1)
	fmt.Println("==========================")
	printBook(&book2)

	fmt.Println("==========================")
	updateBook(&book2, "update_author")
	printBook(&book2)

}
