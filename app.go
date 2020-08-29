package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"

	_ "github.com/bin16/URLprobe/dotenv"
	"github.com/bin16/URLprobe/ogp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/ogp", func(w http.ResponseWriter, r *http.Request) {
		uq := r.URL.Query()["url"]
		if len(uq) == 0 || uq[0] == "" || !isURL(uq[0]) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("url is required"))
			return
		}

		n, err := getHTML(uq[0])
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

func isURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	// ---- Schema ----
	if !strings.HasPrefix(u.Scheme, "http") {
		return false
	}

	// ---- Host ----
	if u.Hostname() == "localhost" {
		return false
	}
	if ip := net.ParseIP(u.Hostname()); ip != nil {
		return false
	}
	// example.com:8080 example.com
	if u.Host != u.Hostname() {
		return false
	}

	return true
}
