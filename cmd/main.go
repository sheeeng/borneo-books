package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sheeeng/borneo-books/pkg/routes"

	// Use the blank identifier as explicit package name to import a package solely for its side-effects (initialization).
	// https://golang.org/ref/spec#Import_declarations
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:65535", r))
}
