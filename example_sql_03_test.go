package databasex

import (
	"context"
	"log"
)

func ExampleNewSimpleSQL_insertTrans() {

	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

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

	// start transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	sqlOp := NewSimpleSQL(postgres)

	// prepare data
	student := Student{"C001", "junjun", 6, 1}

	// create model
	model := NewSimpleModel("tb_student", student)

	// insert data into table
	if err = sqlOp.InsertTrans(context.Background(), tx, model); err != nil {
		log.Fatalln(err.Error())
	}

	tx.Commit() // or tx.Rollback()
}

func ExampleNewSimpleSQL_deleteTrans() {

	// connect to postgresql with database db_belajar_golang
	postgres, err := NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// get database connection
	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// start transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// create sql operatio object
	sqlOp := NewSimpleSQL(postgres)

	// create model we want to work with
	model := NewSimpleModel("tb_student", nil)

	// delete record for student whose name is agus
	if _, err = sqlOp.DeleteTrans(context.Background(), tx, model, "name='agus'"); err != nil {
		log.Fatalln(err.Error())
	}

	// delete all records
	if _, err = sqlOp.DeleteTrans(context.Background(), tx, model, ""); err != nil {
		log.Fatalln(err.Error())
	}

	tx.Commit() // or 	tx.Rollback()

}

func ExampleNewSimpleSQL_updateTrans() {

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

	// start transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	sqlOp := NewSimpleSQL(postgres)

	const (
		cName  = "eko win"
		cAge   = 8
		cGrade = 3
	)

	var keypair = struct {
		Name  string
		Age   int
		Grade int
	}{cName, cAge, cGrade}

	// create model
	model := NewSimpleModel("tb_student", keypair)

	if _, err = sqlOp.UpdateTrans(context.Background(), tx, model, "id='C001'"); err != nil {
		log.Fatalln(err.Error())
	}

	tx.Commit() // or 	tx.Rollback()

}
