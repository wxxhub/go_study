package main

import "fmt"

type Phone interface {
	call()
}

type MeizuPhone struct {

}

type IPhone struct {
	
}

func (m MeizuPhone)call() {
	fmt.Println("I am Meizu, I can call you!")
}

func (m IPhone)call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main()  {
	var phone Phone

	phone = new(MeizuPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}