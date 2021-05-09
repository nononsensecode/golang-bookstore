package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"nononsensecode.com/book/internal/config"
	"nononsensecode.com/book/internal/models"
)

var cfg *config.Config

func main() {
	cfg = config.GetConfig()
	models.InitDB(cfg)

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	repo, err := models.NewBookRepository(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	books, err := repo.FindAll()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range books {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}