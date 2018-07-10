package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("/(\\d*)")
	m := re.FindStringSubmatch(r.URL.Path)
	if 1 == len(m) {
		fmt.Fprintln(w, "XXX: Go to Webkit.org")
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
