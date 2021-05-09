package models

import (
	"log"
	"os"
)

func initSqlite() {
	log.Println("Initialising database in sqlite")
	createTblStmt := `CREATE TABLE books (
			isbn TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			price REAL NOT NULL
		)`
	_, err := db.Exec(createTblStmt)
	if err == nil {
		log.Println("Inserting data to the books database")
		var books = []Book{
			{"978-1503261969", "Emma", "Jayne Austen", 9.44},
			{"978-1505255607", "The Time Machine", "H. G. Wells", 5.99},
			{"978-1503379640", "The Prince", "Niccolò Machiavelli", 6.99},
		}
		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}

		stmt, err := tx.Prepare("INSERT INTO books (isbn, title, author, price) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}

		for _, book := range books {
			_, err = stmt.Exec(book.Isbn, book.Title, book.Author, book.Price)
			if err != nil {
				log.Println(err)
				tx.Rollback()
				os.Exit(2)
			}
		}

		stmt.Close()
		tx.Commit()
	}
}

type SqliteBookRepository struct {}

func (repo SqliteBookRepository) FindAll() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return books, nil
}