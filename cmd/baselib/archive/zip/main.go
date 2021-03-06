package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		//
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("out/out.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
