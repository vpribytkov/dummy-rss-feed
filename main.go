package main

import (
	"fmt"
	"net/http"
	"os"
)

func serveRSS(w http.ResponseWriter, req *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		errMsg := fmt.Sprintf("Can't serve rss: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))
		return
	}

	rss, err := os.ReadFile(fmt.Sprintf("%s/sample_rss.xml", pwd))
	if err != nil {
		errMsg := fmt.Sprintf("Can't serve rss: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))
		return
	}

	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(rss)
}

func main() {
	http.HandleFunc("/rss", serveRSS)

	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
