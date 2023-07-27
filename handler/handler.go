package handler

import (
	"golangweb/entity"
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

	// dynamic data, using map to be parse to HTML
	// data := map[string]string{
	// 	"title":   "Golang Web",
	// 	"content": "Building the backend . . .",
	// }

	// store data to struct for parsing to HTML
	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 9}

	// slice of struct
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 9},
		{ID: 1, Name: "Xpander", Price: 270000000, Stock: 9},
		{ID: 1, Name: "Pajero Sport", Price: 500000000, Stock: 99},
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
	// Store data to map for parsing HTML
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

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("It was GET"))
	case "POST":
		w.Write([]byte("It was POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// load HTML file
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		// Render the template with the data
		tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		// load HTML file
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
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

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
}
