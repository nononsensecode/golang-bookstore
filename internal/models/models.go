package models

type Book struct {
	Isbn 	string
	Title 	string
	Author 	string
	Price 	float32
}

type BookRepository interface {
	FindAll() ([]Book, error)
}