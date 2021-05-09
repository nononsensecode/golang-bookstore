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

func (repo PostgresqlBookRepository) Save(book *Book) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO books (isbn, title, author, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		stmt.Close()
		tx.Rollback()
		return err
	}

	stmt.Close()
	tx.Commit()
	return nil
}