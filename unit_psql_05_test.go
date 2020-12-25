package databasex

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func psqlCreateTableTime01(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add timestamp field to tabel tb_student in db_belajar_golang database using postgresql")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at timestamp)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func psqlCreateTableTime02(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add timestamp with time zone field to tabel tb_student in db_belajar_golang database using postgresql")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at timestamp with time zone)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func psqlCreateTableTime03(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add time field to tabel tb_student in db_belajar_golang database using postgresql")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at time)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func psqlCreateTableTime04(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add time field with time zone to tabel tb_student in db_belajar_golang database using postgresql")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at time with time zone)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func psqlCreateTableTime05(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add date field to tabel tb_student in db_belajar_golang database using postgresql")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at date)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func psqlInsertOneRecord(t *testing.T) (dbData []tStudentTest, prmData tStudentTest, err error) {

	t.Logf("create connection to database server")

	currDb, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	//myStrTime := time.Now().Format("2006-01-02 15:04:05")
	//student := tStudentTest{"C001", "junjun", 6, 1, myStrTime}
	student := tStudentTest{"C001", "junjun", 6, 1, time.Now()}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read table")

	students := make([]tStudentTest, 0)
	model.SetNewData(tStudentTest{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &students); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(students) < 1 {
		t.Errorf("adding one data fail")
		return
	}

	//fmt.Println(data[0].CreatedAt.Day())
	//fmt.Println(student.CreatedAt.Day())
	//fmt.Println(data[0].CreatedAt)
	//fmt.Println(student.CreatedAt)

	if students[0].ID != student.ID || students[0].Name != student.Name || students[0].Age != student.Age ||
		students[0].Grade != student.Grade { //|| isTimeDifferent(data[0].CreatedAt, student.CreatedAt) {

		t.Errorf("data is different")
	}

	log.Println("created_at (db) : ", students[0].CreatedAt) //.Format("2006-01-02 15:03:04"))
	log.Println("created_at (var) : ", student.CreatedAt)    //.Format("2006-01-02 15:03:04"))

	//if !strings.Contains(student.CreatedAt, students[0].CreatedAt) {
	//	t.Errorf("data is different")
	//}

	return students, student, nil
}

func Test_Psql_datetime_01(t *testing.T) {

	log.Println(t.Name())

	psqlCreateTableTime01(t)

	//psqlInsertOneRecord(t)

	students, student, err := psqlInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date and Time is different\n")
	}
}

func Test_Psql_datetime_02(t *testing.T) {

	log.Println(t.Name())

	psqlCreateTableTime02(t)

	//psqlInsertOneRecord(t)

	students, student, err := psqlInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date and Time is different\n")
	}

}

func Test_Psql_datetime_03(t *testing.T) {

	log.Println(t.Name())

	psqlCreateTableTime03(t)

	//psqlInsertOneRecord(t)

	students, student, err := psqlInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Time is different\n")
	}

}

func Test_Psql_datetime_04(t *testing.T) {

	log.Println(t.Name())

	psqlCreateTableTime04(t)

	//psqlInsertOneRecord(t)

	students, student, err := psqlInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date and Time is different\n")
	}

}

func Test_Psql_datetime_05(t *testing.T) {

	log.Println(t.Name())

	psqlCreateTableTime05(t)

	//psqlInsertOneRecord(t)

	students, student, err := psqlInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date is different\n")
	}

}
