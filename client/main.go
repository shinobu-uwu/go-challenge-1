package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Status code not OK. Got: ", resp.StatusCode)
	}

	defer resp.Body.Close()
	price, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("cotacao.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	fileContent := fmt.Sprintf("Dólar: %v", string(price))
	_, err = f.WriteString(fileContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Arquivo criado com sucesso!")
}
