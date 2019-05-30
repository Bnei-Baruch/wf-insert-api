package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"flag"
)

func (a *App) uploadHandler(w http.ResponseWriter, r *http.Request) {

	flag.Parse()
	var s props

	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("r.FormFile(: %v \n", err)
		return
	}

	filename := header.Filename
	filepath := *storage + filename
	mimetype := header.Header.Get("Content-Type")

	s.Filename = filename
	s.Mimetype = mimetype

	defer file.Close()

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file")
		return
	}

	defer file.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	err = s.fileProps(filepath)
	if err != nil {
		return
	}

	respondWithJSON(w, http.StatusOK, s)
}
