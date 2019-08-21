package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func SqlLiteExampleTwo() {
	db, err := sql.Open("sqlite3", "books.sqlite")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	getList(db)

}

func getList(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM book")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	books := make([]*book, 0)

	for rows.Next() {
		row := new(book)
		err := rows.Scan(&row.name, &row.author)

		if err != nil {
			log.Fatal(err)
		}
		books = append(books, row)
	}

	for i, book := range books {
		fmt.Printf("%d %s %s \n", i, book.author, book.name)
	}
}
