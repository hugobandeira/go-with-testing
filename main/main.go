package main

import "fmt"

const prefix = "Fist Testing Hugo "

func Hello(name string) string {
	return prefix + name
}

func main() {
	fmt.Println(Hello("Hugo"))
}
