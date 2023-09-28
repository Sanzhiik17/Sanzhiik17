package main

import (
	"ascii-art-web/asciifunctions"
	"html/template"
	"log"
	"net/http"
)

// type Error struct {
// 	Message string
// 	Status  int
// }
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/window", Posthandler)
	log.Println("Запуск сервера на http://127.0.0.1:4050")
	err := http.ListenAndServe(":4050", mux)
	log.Fatal(err)
}
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errorhandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, 405)
		return
	}
	res, _ := template.ParseFiles("assets/webb.html")
	err := res.Execute(w, nil)
	if err != nil {
		return
	}
}
func Posthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/window" {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}
	res1 := r.FormValue("first")
	for _, v := range res1 {
		if v < 0 || v > 127 {
			Errorhandler(w, 400)
			return
		}
	}
	res2 := r.FormValue("second")
	data, err := ascii(res1, res2)
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	html, err := template.ParseFiles("assets/webb.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		Errorhandler(w, 500)
		return
	}
}
func Errorhandler(w http.ResponseWriter, status int) {
	html, err := template.ParseFiles("assets/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	err = html.Execute(w, http.StatusText(status))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func ascii(s string, font string) (string, error) {
	res1, err := asciifunctions.Split(s)
	if err != nil {
		return "", err
	}
	res, err := asciifunctions.ASCII(res1, font)
	if err != nil {
		return "", err
	}
	return res, nil
}
