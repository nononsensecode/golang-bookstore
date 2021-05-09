package models

import "log"

type PostgresqlBookRepository struct {}

func (repo PostgresqlBookRepository) FindAll() ([]Book, error) {
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