package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

/*
Material covered
1. Routing
2. Handler Function
3. Query string
4. Package for handler
5. Basic HTML
6. Dynamic data to HTML view
7. Layout page
8. Passing struct to view
*/

func main() {
	// Routing
	mux := http.NewServeMux()

	// Calling for Handler Function
	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProdutHandler) // query string case

	// log
	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

/*
Material covered
1. Routing
2. Handler Function
3. Query string
4. Package for handler
5. Basic HTML
6. Dynamic data to HTML view
7. Layout page
8. Passing struct to view
*/

func main() {
	// Routing
	mux := http.NewServeMux()

	// Calling for Handler Function
	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProdutHandler) // query string case

	// log
	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
