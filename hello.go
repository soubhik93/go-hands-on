package main

import "fmt"

func main() {
	fmt.Print(hello("Ali"))

}

func hello(name string) string {
	return "Hello " + name
}
