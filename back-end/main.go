package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Print("HIHIIII")
	log.Print("SADSAD")
	//fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}

func registerHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}

func main() {
	//http.HandleFunc("/", loginHandler)
	//http.HandleFunc("/html-front-end/html/register.html", registerHandler)
	//log.Fatal(http.ListenAndServe("127.0.0.1:5500", nil))
}
