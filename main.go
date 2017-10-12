package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	// The working directory
	homedir string = os.Getenv("HOME")
	// Folder to store files uploaded
	fileStorePath string = path.Join(homedir, "file")
	// Path of web resources
	webResourcePath = path.Join(homedir, "workspace/sdmicroweb")
	// View path
	viewpath string = path.Join(webResourcePath, "view")
)

func index(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join(viewpath, "index.html")
	fmt.Println("file path: ", filePath)
	http.ServeFile(w, r, filePath)
}

func main() {
	port := "6060"

	http.HandleFunc("/", index)

	log.Println("Shendu DevOps web server running on port => ", port)

	http.ListenAndServe(":"+port, nil)
}
