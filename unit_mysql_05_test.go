package databasex

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func mySQLCreateTableTime01(t *testing.T) {

	//log.Println(t.Name())
	log.Println("testing : add datetime field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at datetime)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

/*
func mySQLCreateTableTime02(t *testing.T) {

	//log.Println(t.Name())

	log.Println("testing : add smalldatetime field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int, created_at smalldatetime)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}
*/
func mySQLCreateTableTime03(t *testing.T) {

	//log.Println(t.Name())

	log.Println("testing : add date field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int" +
		",created_at date)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func mySQLCreateTableTime04(t *testing.T) {

	//log.Println(t.Name())

	log.Println("testing : add time field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int" +
		",created_at time)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

/*
func mySQLCreateTableTime05(t *testing.T) {

	//log.Println(t.Name())

	log.Println("testing : add datetimeoffset field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int" +
		",created_at datetimeoffset)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}
*/
func mySQLCreateTableTime06(t *testing.T) {

	//log.Println(t.Name())

	log.Println("testing : add timestamp field to tabel tb_student in db_belajar_golang database using mysql")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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

	cmdStr = "create table if not exists tb_student(id varchar(5), name varchar(255), age int, grade int" +
		",created_at timestamp)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func mySQLInsertOneRecord(t *testing.T) (dbData []tStudentTest, prmData tStudentTest, err error) {

	t.Logf("create connection to database server")

	currDb, err := NewMysql(mysqlUsernameTest, mysqlPasswordTest, mysqlHostTest, mysqlPortTest, mysqlDbTest,
		mysqlMaxConnectionsTest, mysqlMaxIdleTest)
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
	cTime, _ := time.Parse(layOutDateTime, time.Now().Format(layOutDateTime))
	student := tStudentTest{"C001", "junjun", 6, 1, cTime}
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

func Test_MySql_datetime_01(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime01(t)

	students, student, err := mySQLInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date and time are different\n")
	}

}

/*
func Test_MySql_datetime_02(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime02(t)

	 mySQLInsertOneRecord(t)
}
*/

func Test_MySql_datetime_03(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime03(t)

	students, student, err := mySQLInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Date is different\n")
	}

}

func Test_MySql_datetime_04(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime04(t)

	students, student, err := mySQLInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Time is different\n")
	}

}

/*
func Test_MySql_datetime_05(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime05(t)

	 mySQLInsertOneRecord(t)

	}
*/

func Test_MySql_datetime_06(t *testing.T) {

	log.Println(t.Name())

	mySQLCreateTableTime06(t)

	students, student, err := mySQLInsertOneRecord(t)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if isDateTimeDifferent(students[0].CreatedAt, student.CreatedAt) {
		t.Error("Time is different\n")
	}

}
