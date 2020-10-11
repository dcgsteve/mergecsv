package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".csv" {
			// Process each CSV
			fmt.Println(f.Name())
		}
	}
}
