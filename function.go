package main

import "fmt"

func testArgs(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func printType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
			case string:
				fmt.Println(arg, "is a string value.")
			case float64:
				fmt.Println(arg, "is a float64 value.")
			case rune:
				fmt.Println(arg, "is a rune value.")
			case bool:
				fmt.Println(arg, "is a bool value.")
			default:
				fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	testArgs(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("")

	var v1 int = 1
	var v2 int64 = 55
	var v3 string = "hello world!"
	var v4 float64 = 0.0
	var v5 bool = true
	var v6 rune = '6'
	var v7 int8 = 25

	printType(v1, v2, v3, v4, v5, v6, v7)
	fmt.Println("")

	// 匿名函数	
	func(input string) {
		fmt.Println("input:", input)
	}("Hello")
	fmt.Println("")

	j := 5
	
	fmt.Println("create a")
	a := func()(func())  {
		i := 10
		fmt.Printf("func i, j : %d, %d\n", i, j)
		return func()  {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}
	fmt.Println("")

	fmt.Println("create b")
	b := func()(func())  {
		i := 10
		fmt.Printf("func i, j : %d, %d\n", i, j)
		return func()  {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	fmt.Println("")

	fmt.Println("a():")
	a()
	fmt.Println("")
	
	fmt.Println("b():")
	b()
	fmt.Println("")

	j *= 2

	fmt.Println("a():")
	a()
	fmt.Println("")

	fmt.Println("b():")
	b()
	fmt.Println("")
}