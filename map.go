package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo = make(map[string] PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 1502,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 6512,..."}

	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID", person.ID)
	} else {
		fmt.Println("Did not find person with ID 12345.")
	}
}