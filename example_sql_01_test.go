package databasex

import (
	"context"
	"fmt"
	"log"
)

func ExampleNewSimpleSQL_insertDb() {

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

	sqlOp := NewSimpleSQL(postgres)

	// prepare data
	student := Student{"C001", "junjun", 6, 1}

	// create model
	model := NewSimpleModel("tb_student", student)

	// insert data into table
	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		log.Fatalln(err.Error())
	}

}

func ExampleNewSimpleSQL_deleteDb() {

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

	// create sql operatio object
	sqlOp := NewSimpleSQL(postgres)

	// create model we want to work with
	model := NewSimpleModel("tb_student", nil)

	// delete record for student whose name is agus
	if _, err = sqlOp.DeleteDb(context.Background(), model, "name='agus'"); err != nil {
		log.Fatalln(err.Error())
	}

	// delete all records
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		log.Fatalln(err.Error())
	}

}

func ExampleNewSimpleSQL_updateDb() {

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

	if _, err = sqlOp.UpdateDb(context.Background(), model, "id='C001'"); err != nil {
		log.Fatalln(err.Error())
	}

}

func ExampleNewSimpleSQL_selectDb() {

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

	sqlOp := NewSimpleSQL(postgres)

	// create model
	model := NewSimpleModel("tb_student", Student{})

	data := make([]Student, 0)

	// select record for ID = C001
	if err = sqlOp.SelectDb(context.Background(), model, "ID='C001'", &data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(data)

	data = make([]Student, 0)

	// select all records
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(data)

}
