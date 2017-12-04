package main

import "fmt"

type Book struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Book
	Book1.author="张三"
	Book1.title="钢铁是怎么炼成的"
	Book1.subject="钢铁"
	Book1.book_id=1
	printBook(Book1)
	printBookZhiZhen(&Book1)
}

func printBook(Book1 Book)  {
	fmt.Printf("boo1.title=%s\n",Book1.title)
	fmt.Printf("book1.author=%s\n",Book1.author)
	fmt.Printf("book1.subject=%s\n",Book1.subject)
	fmt.Printf("book1.book_id=%d\n",Book1.book_id)

}

func printBookZhiZhen(book*Book)  {
	fmt.Printf("指针boo1.title=%s\n",book.title)
	fmt.Printf("指针book1.author=%s\n",book.author)
	fmt.Printf("指针book1.subject=%s\n",book.subject)
	fmt.Printf("指针book1.book_id=%d\n",book.book_id)

}
