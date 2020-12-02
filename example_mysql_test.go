package databasex

import (
	"fmt"
	"log"
)

func ExampleNewMysql() {
	currDb, err := NewMysql("root", "", "localhost", "3306", "db_belajar_golang",
		0, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	valuesmark, _ := currDb.CreateValuesMark(5)
	fmt.Println(valuesmark)
	//Output:
	// ?,?,?,?,?

}
