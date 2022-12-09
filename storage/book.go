package storage

import (
	"database/sql"

	"app/models"
)

func Create(db *sql.DB, book models.Book) (string, error) {

	var (
		id 		string
		query 	string
	)

	query = `
		INSERT INTO
			book (title, price, author_name, published_year, page, genre)
		VALUES ( $1, $2, $3, $4, $5, $6 )
		RETURNING id
	`
	err := db.QueryRow(
		query,
		book.Title,
		book.Price,
		book.AuthorName,
		book.PublishedYear,
		book.Page,
		book.Genre,	
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById (db *sql.DB, id string) (models.Book, error) {

	var (
		book models.Book
		query string
	)

	query = `
		SELECT 
			id, 
			title, 
			price, 
			author_name, 
			published_year, 
			page, 
			genre
		FROM
			book
		WHERE id = $1 
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&book.Id,
		&book.Title,
		&book.Price,
		&book.AuthorName,
		&book.PublishedYear,
		&book.Page,
		&book.Genre,
	)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func GetList(db *sql.DB) ([]models.Book, error) {

	var (
		book []models.Book
		query string
	)

	query = `
		SELECT
		id, 
		title, 
		price, 
		author_name, 
		published_year, 
		page, 
		genre
		FROM
			book
	`

	rows, err := db.Query(query)

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		var books models.Book

		err = rows.Scan(
			&books.Id,
			&books.Title,
			&books.Price,
			&books.AuthorName,
			&books.PublishedYear,
			&books.Page, 	
			&books.Genre, 
		)

		if err != nil {
			return []models.Book{}, err
		}

		book =append(book, books) 
	}

	return book, nil 


	// func Update(db *sql.DB, book models.Book) (string, error) {

	// 	var (
	// 		id 		string
	// 		query 	string
	// 	)
	
	// 	query = `
	// 		UPDATE book
	// 			SET title = $2, price = $3 , author_name = $4 , published_year = $5, page = $6, genre = $7
	// 		WHERE id = $1
	// 		RETURNING  	
	// 	`
	// 	_, err = db.Exec(query,)
	
	// 	if err != nil {
	// 		return "", err
	// 	}
	
	// 	return id, nil
	// }
	

}