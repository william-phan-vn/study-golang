package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world!"
	fmt.Println(whatToSay)

	i = 9
	fmt.Println("i is set to", i)

	whatWasSaid, otherThing := saySomeThing()
	fmt.Println("The function returned", whatWasSaid, otherThing)
}

func saySomeThing() (string, string) {
	return "something", "else"
}
