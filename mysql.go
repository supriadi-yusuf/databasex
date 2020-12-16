package databasex

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type mysqlDb struct {
	realDb
}

func (workDb *mysqlDb) GetValueMark(fieldSec int) (valueMark string, err error) {
	return "?", nil
}

// CreateQueryString is method implementing interface
func (workDb *mysqlDb) CreateValuesMark(fieldNum int) (string, error) {

	var arrPrms []string

	for i := 0; i < fieldNum; i++ {
		//arrPrms = append(arrPrms, fmt.Sprintf("$%d", i))
		newValueMark, _ := workDb.GetValueMark(i)
		arrPrms = append(arrPrms, newValueMark)
	}

	return strings.Join(arrPrms, ","), nil
}

// CreateConnection is method
func (workDb *mysqlDb) createConnection(username, password, host, port, dbname string,
	maxConnections, maxIdleConnection int) (*sql.DB, error) {
	//"root:@tcp(127.0.0.1:3306)/db_belajar_golang"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if maxConnections != 0 {
		db.SetMaxOpenConns(maxConnections)
		db.SetMaxIdleConns(maxIdleConnection)
	}

	workDb.db = db

	return db, nil
}

func (workDb *mysqlDb) BeforeScan(structData reflect.Value) []reflect.Value {

	storages := make([]reflect.Value, 0)

	structDataType := structData.Type()
	for i := 0; i < structData.NumField(); i++ {

		field := structData.Field(i)
		if field.Type().Kind() == reflect.Struct {

			fieldType := structDataType.Field(i)
			tagName := fieldType.Tag.Get(fieldTbl)
			if tagName == "" { // check if it has tag
				newStorages := workDb.BeforeScan(field) //recursive
				storages = append(storages, newStorages...)
				continue
			}
		}

		fieldAddres := field.Addr()
		if field.Type().Name() == "Time" {
			//newField := reflect.New(reflect.TypeOf([]byte(""))).Elem()
			newField := reflect.New(reflect.TypeOf((*interface{})(nil)).Elem()).Elem()
			fieldAddres = newField.Addr()
			//fmt.Println(field.Type().Name())
		}

		storages = append(storages, fieldAddres)
	}

	return storages
}

func (workDb *mysqlDb) AfterScan(structData reflect.Value, prms []reflect.Value) {

	structDataType := structData.Type()
	for i := 0; i < structData.NumField(); i++ {

		field := structData.Field(i)
		if field.Type().Kind() == reflect.Struct {

			fieldType := structDataType.Field(i)
			tagName := fieldType.Tag.Get(fieldTbl)
			if tagName == "" { // check if it has tag
				workDb.AfterScan(field, prms) //recursive
				continue
			}
		}

		if field.Type().Name() == "Time" {

			stime := fmt.Sprintf("%s", prms[i].Elem().Interface())
			//fmt.Println(stime)

			newTime, _ := time.Parse("2006-01-02 15:04:05", stime)
			//fmt.Println(newTime)
			//fmt.Println(field.CanSet())
			field.Set(reflect.ValueOf(newTime))
		}

	}
}

// NewMysql is a function to connect with mysql database.
//
// This function has several input parameters :
//
// - username is username of database we want to access
//
// - password is password of username
//
// - host is location where mysql lives
//
// - port is database port
//
// - dbname is name of database
//
func NewMysql(username, password, host, port, dbname string,
	maxConnections, maxIdleConnection int) (db IDatabase, err error) {

	var workDb mysqlDb

	_, err = workDb.createConnection(username, password, host, port, dbname,
		maxConnections, maxIdleConnection)
	if err != nil {
		return nil, err
	}

	return &workDb, nil
}
