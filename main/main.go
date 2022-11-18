package main

import (
	"fmt"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "OsbornCollins"

	fmt.Fprintf(response, "A little bit about Osborn Collins...")
	fmt.Println("Enpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutMe", aboutMe)
}
