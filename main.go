package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <product_name>", os.Args[0])
	}

	productName := strings.Join(os.Args[1:], "-")
	url := "https://lista.mercadolivre.com.br/"

	resp, err := fetchURL(url + productName)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	products, err := parseHTML(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	for i, p := range products {
		if i > 3 {
			break
		}

		fmt.Println("-----")
		fmt.Printf("Título: %v\n", p.Title)
		fmt.Printf("Preço: %v\n", p.Price)
		if p.Discount == "" {
			fmt.Println("Sem desconto")
		} else {
			fmt.Printf("Desconto: %v\n", p.Discount)
		}
		fmt.Printf("Link: %v\n", p.Link)
	}

	err = saveHTMLToFile(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}
}
