package databasex

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type postgresDb struct {
	realDb
}

func (workDb *postgresDb) GetValueMark(fieldSec int) (valueMark string, err error) {
	return fmt.Sprintf("$%d", fieldSec+1), nil
}

// CreateQueryString is method implementing interface
func (workDb *postgresDb) CreateValuesMark(fieldNum int) (string, error) {

	var arrPrms []string

	for i := 0; i < fieldNum; i++ {
		newValueMark, _ := workDb.GetValueMark(i)
		arrPrms = append(arrPrms, newValueMark)
		//arrPrms = append(arrPrms, "?")
	}

	return strings.Join(arrPrms, ","), nil
}

// CreateConnection is method
func (workDb *postgresDb) createConnection(username, password, host, port, dbname, other string,
	maxConnections, maxIdleConnection int) (*sql.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s", username, password, host, port, dbname, other)

	db, err := sql.Open("postgres", connString)
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

func (workDb *postgresDb) BeforeScan(structData reflect.Value) []reflect.Value {
	return generateStorage(structData)
}

func (workDb *postgresDb) AfterScan(structData reflect.Value, prms []reflect.Value) {

}

// NewPostgre is a function to connect with postgresql database.
//
// This function has several input parameters :
//
// - username is username of database we want to access
//
// - password is password of username
//
// - host is location where postgresql lives
//
// - port is database port
//
// - dbname is name of database
//
// - other is additional parameter if we need it. for example : sslmode=disable
func NewPostgre(username, password, host, port, dbname, other string,
	maxConnections, maxIdleConnection int) (db IDatabase, err error) {

	var workDb postgresDb

	_, err = workDb.createConnection(username, password, host, port, dbname, other,
		maxConnections, maxIdleConnection)
	if err != nil {
		return nil, err
	}

	return &workDb, nil
}
