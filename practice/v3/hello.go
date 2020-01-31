package main

import "fmt"

func main() {
	fmt.Println(Hello("blah"))
}

func Hello(name string) string {
	return "Hello, " + name
}
