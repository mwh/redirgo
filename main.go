package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"path"
)

var rootPath string

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "../") {
		w.WriteHeader(403)
		fmt.Fprintf(w, "403 permission denied")
		return
	}
	itemPath := path.Join(rootPath, r.URL.Path)
	content, err := ioutil.ReadFile(itemPath)
	if err != nil {
		content, err = ioutil.ReadFile(path.Join(itemPath, "/.index"))
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "404 not found")
			return
		}
	}
	dest := string(content)
	dest = strings.TrimSpace(dest)
	if !strings.HasPrefix(dest, "https:") && !strings.HasPrefix(dest, "http:") {
		w.WriteHeader(404)
		fmt.Fprintf(w, "404 not found")
		return
	}
	w.Header().Add("Location", dest)
	w.WriteHeader(302)
	fmt.Fprintf(w, "redirecting to %s", dest)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("must provide two arguments PATH and PORT")
	}
	rootPath = os.Args[1]
	port := os.Args[2]
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
