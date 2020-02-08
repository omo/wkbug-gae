package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("/(\\d*)")
	m := re.FindStringSubmatch(r.URL.Path)
	if 1 == len(m) {
		http.Redirect(w, r, "https://www.webkit.org/", http.StatusMovedPermanently)
		return
	}

	bugid, err := strconv.Atoi(m[1])
	if err != nil {
		http.Redirect(w, r, "https://www.webkit.org/", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, "https://b.webkit.org/show_bug.cgi?id="+strconv.Itoa(bugid), http.StatusMovedPermanently)
}
