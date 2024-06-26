package main

import (
	urlshort "URL_Shortener"
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://github.com/Nabil-Salah/URL_Shortener",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/Nabil-Salah/URL_Shortener
- path: /urlshort-final
  url: https://https://github.com/Nabil-Salah/URL_Shortener/branches
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

	/*fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)*/
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dsd")
	fmt.Fprintln(w, "Hello, world!")
}
