package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func (book *Books) Subject() string {
	return book.subject
}

func (book *Books) SetSubject(subject string) {
	book.subject = subject
}

func NewBooks(title string, author string, subject string, book_id int) *Books {
	return &Books{title: title, author: author, subject: subject, book_id: book_id}
}

//struct method
func (book *Books) getBook() {
	fmt.Println("getBook:")
	fmt.Printf("titile:%s \nauthor:%s \nsubject:%s \nbook_id:%d\n", book.title, book.author, book.subject, book.book_id)
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
	book2 := Books{"python book2", "fanpf", "python language", 111222}

	//new创建的不需要手动delete， 有垃圾回收机制
	book3 := NewBooks("book3", "fanyf", "javascript", 8888)

	book1.title = "go book1"
	book1.author = "www.runoob.com"
	book1.subject = "Go book"
	book1.book_id = 64782

	printBook(&book1)
	fmt.Println("==========================")
	printBook(&book2)

	fmt.Println("==========================")
	printBook(&book2)
	book2.SetSubject("new subject")
	fmt.Println("after SetSubject:", book2.Subject())

	//call struct method
	book1.getBook()
	book2.getBook()
	book3.getBook()
}
