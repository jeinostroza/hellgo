package main

import (
	"fmt"
	"os"
)

func translate(imput string) string {
	var translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["es"] = "Hola"
	translations["fr"] = "Bonjour"
	return translations[imput]
}

func argsswitch(input []string) string {
	var locale string
	if len(os.Args) == 1 {
		fmt.Printf("Please, select a language(en,es,de)\n")
		fmt.Scanln(&locale)
	} else {
		locale = os.Args[1]
	}
	return locale
}

func main() {
	locale := argsswitch(os.Args)
	output := translate(locale)
	if output == "" {
		output = "Yo"
	}
	fmt.Printf(output + " Go\n")
}
