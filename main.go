package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type bowerrc struct {
	Directory string `json:"directory"`
}

type bowerjson struct {
	Name string `json:"name"`
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mustReadFile(fileName string) []byte {
	file, e := ioutil.ReadFile(fileName)
	must(e)
	return file
}

func main() {
	// init port
	var port = flag.String("p", "8080", "the port to serve")
	flag.Parse()

	// reading `.bowerrc` file if exists to get components directory
	bowerDirectory := "bower_components"
	if _, err := os.Stat(".bowerrc"); err == nil {
		fileContent := mustReadFile(".bowerrc")
		var rc bowerrc
		must(json.Unmarshal(fileContent, &rc))
		bowerDirectory = rc.Directory
	}

	var elementName string
	fileContent := mustReadFile("./bower.json")
	var bower bowerjson
	must(json.Unmarshal(fileContent, &bower))
	elementName = bower.Name

	fmt.Printf("Start serving on port %s\n", *port)
	fmt.Printf("Serving components from %s\n", bowerDirectory)
	fmt.Printf("Files in this directory are available at localhost:%s/%s/...\n", *port, elementName)

	components := "/"

	http.HandleFunc(components, func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		requestedFile := urlPath[len(components):]

		split := strings.Split(requestedFile, "/")
		base := split[0]
		remain := strings.Join(split[1:], "/")

		var serve string
		if base == elementName {
			serve = remain
		} else {
			serve = bowerDirectory + "/" + requestedFile
		}
		http.ServeFile(w, r, serve)
	})

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
