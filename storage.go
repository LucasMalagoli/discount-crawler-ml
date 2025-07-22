package main

import (
	"io"
	"os"
)

func saveHTMLToFile(body io.Reader) error {
	file, err := os.Create("output.html")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	return err
}
