// Package databasex is to help developer for building app with golang especially when need to interact with dbms.
package databasex

import (
	"database/sql"
	"errors"
	"reflect"
	//_ "github.com/lib/pq"
)

// IDatabase is interface related to dbms. This interface has several methods :
//
// - CreateValuesMark(fieldNum int) (valuesMark string, err error)
//   Create values mark in sql statement.
//   If we use sql to insert data into postgresql. we create command :
//   "insert into tb_xxx(field1, field2, field3) values($1,$2,$3)".
//   $1,$2,$3 are values mark on postgresql.
//
//   But if we want to use mysql, our command should be :
//   "insert into tb_xxx(field1, field2, field3) values(?,?,?)"
//   ?,?,? are values mark on mysql.
//
// - GetDbConnection() (dbConn *sql.DB, err error)
//   Get dbms's connection.
//
type IDatabase interface {
	GetValueMark(fieldSec int) (valueMark string, err error)
	CreateValuesMark(fieldNum int) (valuesMark string, err error)
	GetDbConnection() (dbConn *sql.DB, err error)
	BeforeScan(structData reflect.Value) []reflect.Value
	AfterScan(structData reflect.Value, prms []reflect.Value)
}

type realDb struct {
	db *sql.DB
}

// GetConnection is method
func (workDb *realDb) GetDbConnection() (dbConn *sql.DB, err error) {

	if workDb.db == nil {
		return nil, errors.New("Connection does not exist")
	}

	return workDb.db, nil

}
