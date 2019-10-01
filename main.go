package main

import "fmt"

func main() {

	var locale, greeting string

	fmt.Printf("Please, select a language(en,es,de)\n")
	fmt.Scanln(&locale)
	/*
		if locale == "en" {
			greeting = "Hello"
		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yo"
		}
	*/
	switch locale {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "de":
		greeting = "Guten Tag"
	}

	fmt.Printf(greeting + " Go\n")

}
