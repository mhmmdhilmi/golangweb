package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// Handler function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path) //print routing to console

	// prevent unlisted routing redirect to root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// load HTML file
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// dynamic data, for this case we use static example to be parse to HTML
	data := map[string]string{
		"title":   "Golang Web",
		"content": "Building the backend . . .",
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path) //print routing to console
	w.Write([]byte("hello world!"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path) //print routing to console
	w.Write([]byte("Mario from Nintendo"))
}

func ProdutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path) //print routing to console

	// get query string
	id := r.URL.Query().Get("id")

	// handler for id (should be integer)
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// Static view from golang not HTML view
	//w.Write([]byte("Product Page \n"))
	//fmt.Fprintf(w, "Product id : %d", idNumb)

	// change from fmt.Fprintf to HTML view
	// Store data to map for parsing
	data := map[string]interface{}{
		"content": idNumb,
	}

	// load HTML file
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// Render the template with the data
	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

}
