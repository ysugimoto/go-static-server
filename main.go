package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

		requestPath := filepath.Clean(filepath.Join(documentRoot, url.Path))

		if !strings.HasPrefix(requestPath, documentRoot) {
			w.WriteHeader(403)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Forbidden"))
			return
		}

		fmt.Printf("%v\n", requestPath)
		statInfo, statErr := os.Stat(requestPath)
		if statErr != nil {
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%v\n", statErr)))
			return
		}

		if statInfo.IsDir() {
			entries, readDirErr := ioutil.ReadDir(requestPath)
			if readDirErr != nil {
				w.WriteHeader(500)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(fmt.Sprintf("%v\n", readDirErr)))
				return
			}

			entryLinks := make([]string, 0)
			for _, entry := range entries {
				relPath := strings.Replace(filepath.Join(requestPath, entry.Name()), documentRoot, "" , 1)
				var entryType string
				if entry.IsDir() {
					entryType = "D"
				} else {
					entryType = "F"
				}
				entryLinks = append(entryLinks, fmt.Sprintf("<li>[<span>%s</span>] <a href='%s'>%s</a></li>", entryType, relPath, entry.Name()))
			}

			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Directory listing of %s</title>
		<style>span { min-width: 1em; display: inline-block; text-align: center; }</style>
	</head>
	<body>
		<p>Directory listing of %s</p>
		<ol>
			%s
		</ol>
	</body>
</html>`, url.Path, url.Path, strings.Join(entryLinks, "\n"))))
			return
		}

		buf := bytes.NewBuffer(nil)
		file, _ := os.Open(requestPath)
		io.Copy(buf, file)

		content := buf.Bytes()
		w.Header().Set("Content-Type", http.DetectContentType(content[:512]))
		w.WriteHeader(200)
		defer file.Close()

		w.Write(content)
	})
	http.ListenAndServe(port, nil)
}
