package main

import (
	"fmt"
	"os"
)

func main() {

	var locale string
	var translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["es"] = "Hola"
	translations["fr"] = "Bonjour"

	if len(os.Args) == 1 {
		fmt.Printf("Please, select a language(en,es,de)\n")
		fmt.Scanln(&locale)
	} else {
		locale = os.Args[1]
	}

	output := translations[locale]
	if output == "" {
		output = "Yo"
	}

	fmt.Printf(output + " Go\n")

}
