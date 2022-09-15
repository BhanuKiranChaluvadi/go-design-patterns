package main

import "fmt"

// Iterate using a callback function
func main() {
	// TODO: use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// TODO: Use IterateBooks to iterate via anonymous function
	// lib.IterateBooks(func(b Book) error {
	// fmt.Println("Book author:", b.Author)
	// return nil
	// })

	// TODO: create a BookIterator
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}

}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
