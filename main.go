package main

import "fmt"

func main() {
	//var whatToSay string
	//whatToSay = "Hello, world again!"

	whatToSay := "hello world, again!"

	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
