package databasex

import (
	"fmt"
	"log"
)

func ExampleNewPostgre() {
	postgres, err := NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	valuesmark, _ := postgres.CreateValuesMark(5)
	fmt.Println(valuesmark)
	//Output:
	// $1,$2,$3,$4,$5

}
