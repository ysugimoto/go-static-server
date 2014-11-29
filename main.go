package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
)

func main() {
	var (
		documentRoot string
		portNumber   int
	)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error, cannot get working directory: %v", err)
		os.Exit(1)
	}

	flag.IntVar(&portNumber, "p", 8888, "Listen port number")
	flag.StringVar(&documentRoot, "d", cwd, "Document Root")
	flag.Parse()

	fmt.Printf("DocumentRoot is %s\n", documentRoot)
	fmt.Printf("Listen 0.0.0.0:%d...\n", portNumber)
	port := ":" + fmt.Sprint(portNumber)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL

		matched, err := regexp.Match("/$", []byte(url.Path))
		if err != nil {
			fmt.Println("Path find error")
			return
		}

		if matched {
			url.Path += "index.html"
		}

		requestPath := documentRoot + url.Path

		fmt.Printf("%v\n", requestPath)
		_, statErr := os.Stat(requestPath)
		if statErr != nil {
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%v\n", statErr)))
			return
		}

		ext := path.Ext(requestPath)
		switch ext {
		case ".txt":
			w.Header().Set("Content-Type", "text/plain")
		case ".html":
		case ".htm":
			w.Header().Set("Content-Type", "text/html")
		case ".gif":
			w.Header().Set("Content-Type", "image/gif")
		case ".jpg":
		case ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".json":
			w.Header().Set("Content-Type", "application/json")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		default:
			w.Header().Set("Content-Type", "application/octet-stream")
		}
		w.WriteHeader(200)

		buf := bytes.NewBuffer(nil)
		file, _ := os.Open(requestPath)
		io.Copy(buf, file)
		defer file.Close()

		w.Write(buf.Bytes())
	})
	http.ListenAndServe(port, nil)
}
