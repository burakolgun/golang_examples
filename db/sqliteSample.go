package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type book struct {
	name string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "books.sqlite")

	var book book

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS book (name text, author text);`)
	db.Exec(`INSERT INTO book VALUES
		('Design Pattern in PHP and Laravel', 'Kelt Dockins'),
		('Apprenticeship Patterns', 'Dave H. Hoover & Adawale Oshineye')
	`)

	db.QueryRow("SELECT * FROM book;").Scan(&book.name, &book.author)
	fmt.Println(book)
}
