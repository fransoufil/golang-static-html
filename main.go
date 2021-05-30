package main

import (
	"log"
	"net/http"
)

func getStaticHtml() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Servidor está rodando na porta 1337...")
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	getStaticHtml()

}
