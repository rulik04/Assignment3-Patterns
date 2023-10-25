package main

import (
	"fmt"
	"sync"
)

type President struct {
	Name string
}

var presidentInstance *President
var once sync.Once

func GetPresident() *President {
	once.Do(func() {
		presidentInstance = &President{Name: "Nursultan Nazarbayev"}
	})
	return presidentInstance
}

func main() {
	president1 := GetPresident()
	president1.Name = "Qasym-Jomart Toqayev"

	president2 := GetPresident()

	fmt.Printf("President 1: %s\n", president1.Name)
	fmt.Printf("President 2: %s\n", president2.Name)
}
