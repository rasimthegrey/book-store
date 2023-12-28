package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

type Author struct {
	AuthorID   int64
	AuthorName string
}

type Category struct {
	CategoryID   int64
	CategoryName string
}

type Book struct {
	BookID       int64
	BookName     string
	ISBN         string
	BookCategory Category
	BookAuthor   Author
}

func ConnectDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "LibrarySystem",
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected")
	return db, nil
}

func QueryBookByID(id int) (Book, error) {
	db, err := ConnectDB()
	if err != nil {
		return Book{}, err
	}
	defer db.Close()
	var book Book

	row := db.QueryRow("SELECT Book.BookName, Book.ISBN, Category.CategoryName, Author.AuthorName FROM ((Book INNER JOIN Category ON Category.CategoryID=Book.CategoryID) INNER JOIN Author ON Author.AuthorID=Book.AuthorID) WHERE Book.BookID = ?", id)
	if err := row.Scan(&book.BookName, &book.ISBN, &book.BookCategory.CategoryName, &book.BookAuthor.AuthorName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return book, fmt.Errorf("%d, no such book", id)
		}
		return book, fmt.Errorf("GetBookByID %d: %v", id, err)
	}
	return book, nil
}

func QueryAllBooks() ([]Book, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var books []Book

	rows, err := db.Query("SELECT Book.BookName, Book.ISBN, Category.CategoryName, Author.AuthorName FROM ((Book INNER JOIN Category ON Category.CategoryID=Book.CategoryID) INNER JOIN Author ON Author.AuthorID=Book.AuthorID)")
	if err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookName, &book.ISBN, &book.BookCategory.CategoryName, &book.BookAuthor.AuthorName); err != nil {
			return nil, fmt.Errorf("error: %q", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	return books, nil
}

func QueryCategories() ([]Category, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var categories []Category
	rows, err := db.Query("SELECT CategoryName FROM Category")
	if err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryName); err != nil {
			return nil, fmt.Errorf("error: %q", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	return categories, nil
}

func QueryAuthors() ([]Author, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var authors []Author
	rows, err := db.Query("SELECT AuthorName FROM Author")
	if err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.AuthorName); err != nil {
			return nil, fmt.Errorf("error: %q", err)
		}
		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %q", err)
	}
	return authors, nil
}
