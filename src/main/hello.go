package main

import "fmt"

var age = 15

func getAge() int {
	return age
}

func setAge(a int) {
	age = a
}

func main() {
	setAge(100)
	fmt.Println("my age is", getAge())
}
