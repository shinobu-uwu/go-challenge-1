package routes

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shinobu-uwu/go-challenge-1/server/models"
)

const ENDPOINT = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func Cotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ENDPOINT, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed to create request: ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed sending the request", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed to read response body", err)
		return
	}

	price := models.Response{}
	err = json.Unmarshal(body, &price)

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed to unmarshal json", err)
		return
	}

	err = price.USDBRL.Save()

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed to insert into database", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(price.USDBRL.Price))

	if err != nil {
		http.Error(w, "Could not fulfill the request", http.StatusInternalServerError)
		log.Default().Println("Failed to copy response body to response", err)
	}
}
