package main

import "fmt"

func main() {
	my_slice := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(my_slice); i++ {
		fmt.Printf("1. my_slice[%d] = %d\n", i, my_slice[i])
	}

	fmt.Println("")

	my_slice = append(my_slice, 6, 7, 8)
	for i, value := range my_slice {
		fmt.Printf("2. my_slice[%d] = %d\n", i, value)
	}
	
	fmt.Println("")

	old_slice := my_slice[2:4]

	new_slice := append(my_slice, old_slice...)

	for i, value := range old_slice {
		fmt.Printf("old_slice[%d] = %d\n", i, value)
	}

	fmt.Println("")

	for i, value := range new_slice {
		fmt.Printf("new_slice[%d] = %d\n", i, value)
	}

	fmt.Println("")

	size := cap(new_slice) // new_slice size

	cp_slice := make([]int, size - 2)

	copy(cp_slice, new_slice)

	for i, value := range cp_slice {
		fmt.Printf("cp_slice[%d] = %d\n", i, value)
	}

	fmt.Println("")
}