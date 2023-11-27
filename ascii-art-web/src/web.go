package src

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Text   string
	Banner string
	Ascii  string
}

var data PageData

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		NotFoundHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)

	data := PageData{
		Text:   "Hello, World!",
		Banner: "standard",
		Ascii:  "Your ASCII art will be displayed here!",
	}

	tmpl, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data.Text = r.FormValue("text")
		data.Banner = r.FormValue("banner")
		data.Ascii = GenerateAscii(data.Text, data.Banner)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
