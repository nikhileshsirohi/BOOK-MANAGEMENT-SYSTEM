package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nikhileshsirohi/BOOK-MANAGEMENT-SYSTEM/pkg/routes"
)

func main() {
	fmt.Println("WELCOME TO BOOK MANAGEMENT SYSTEM")
	r := mux.NewRouter()
	fmt.Println("Server Start at Port: 9010")
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
