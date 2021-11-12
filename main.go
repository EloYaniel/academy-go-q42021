package main

import (
	"log"
	"net/http"

	app "github.com/EloYaniel/academy-go-q42021/app"
)

func main() {
	r := app.InitApp()
	log.Fatal(http.ListenAndServe(":8080", r))
}
