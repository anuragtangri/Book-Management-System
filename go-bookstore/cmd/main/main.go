package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anuragtangri/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routes.RegisterBookStoreRoutes(r)
	r.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
