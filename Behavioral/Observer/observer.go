package main

import "fmt"

// Define an interface for an observer class
type observer interface {
	onUpdate(data string)
}

// Our data observer will have a name
type DataListener struct {
	Name string
}

// TODO: TO conform to the interface, it must have an onUpdate function
func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", dl.Name, "got data change:", data)
}
