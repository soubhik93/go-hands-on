package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello "
	spanishHello = "Hola "
	frenchHello  = "Bonjour "
	defaultName  = "World"
)

func main() {
	fmt.Print(hello("Ali", "English"))

}

func hello(name string, language string) string {

	if name == "" {
		name = defaultName
	}

	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello

	}
	return
}
