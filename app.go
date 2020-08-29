package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"

	_ "github.com/bin16/URLprobe/dotenv"
	"github.com/bin16/URLprobe/ogp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/ogp", func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.Query()["url"][0]
		n, err := getHTML(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		results := ogp.Parse(n)

		data, err := json.Marshal(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", 7200))
		w.Write(data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "2222"
	}

	log.Printf("Server running at %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func getHTML(pageURL string) (*html.Node, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return n, nil
}
