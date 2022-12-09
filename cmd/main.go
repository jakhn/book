package main

import (
	"fmt"
	"log"

	"app/config"
	"app/pkg/db"
	"app/storage"
	"app/models"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	b := models.Book{ 
	Title: "The Cost of Living",
	Price: 299000,
	AuthorName: " Deborah Levy",
	PublishedYear: 2018,
	Page : 400, 
	Genre: "biography",
	}

	id, err := storage.Create(db, b)
	if err != nil {
		log.Fatal("error while creating new book: %v", err)
	}

	fmt.Println(id)



	book, err := storage.GetById(db, id)
	if err != nil {
		log.Fatalf("error while creating new book : %v", err)
	}

	fmt.Println(book)
	
	// bookList, err := storage.GetList(db)
	// if err != nil {
	// 	log.Fatalf("error while creating new book : %v", err)
	// }

	// fmt.Println(bookList)



	// bookList, err := storage.GetList(db)
	// if err != nil {
	// 	log.Fatalf("error while creating new book : %v", err)
	// }

	// fmt.Println(bookList) 


}
