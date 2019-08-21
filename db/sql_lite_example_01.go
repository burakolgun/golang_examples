package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type book struct {
	name   string
	author string
}

func SqlLiteExampleOne() {
	db, err := sql.Open("sqlite3", "books.sqlite")

	var book book

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS book (name text, author text);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO book VALUES
		('Design Pattern in PHP and Laravel', 'Kelt Dockins'),
		('Apprenticeship Patterns', 'Dave H. Hoover & Adawale Oshineye')
	`)

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow("SELECT * FROM book;").Scan(&book.name, &book.author)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book)
}
