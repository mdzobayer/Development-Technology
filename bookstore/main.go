package main

import (
	"bookstore/models"
	"fmt"
	"net/http"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

func main() {
	models.InitDB("root:@tcp(localhost:3306)/bookstore")
	//fmt.Println("After Connection Create")

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //fmt.Println("Before Query")
	// rows, err := db.Query("SELECT * FROM books")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// //fmt.Println("After Query")

	// bks := make([]*Book, 0)
	// for rows.Next() {
	// 	bk := new(Book)
	// 	err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bks = append(bks, bk)
	// }
	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, bk := range bks {
	// 	fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	// }

}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
