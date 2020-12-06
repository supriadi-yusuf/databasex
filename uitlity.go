package databasex

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const fieldTbl = "fieldtbl"

func extractFromModel(model IModel) (tblName string, fields []string, values []interface{}, err error) {
	tblName = model.GetTableName()

	data := model.GetData()

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}

	if dataValue.Kind() != reflect.Struct {
		err = errors.New("model data must be struct type")
		return
	}

	fields, err = getFieldsFromType(dataValue.Type())
	if err != nil {
		return
	}

	values, err = getValuesFromValue(dataValue)
	if err != nil {
		return
	}

	return
}

func getFieldsFromType(dataType reflect.Type) (fields []string, err error) {

	slices := make([]string, 0)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)

		name := field.Name
		tagName := field.Tag.Get(fieldTbl)
		if tagName != "" {
			name = tagName
		}

		if field.Type.Kind() == reflect.Struct {
			if tagName == "" { // check if it has tag

				names, err := getFieldsFromType(field.Type) //recursive
				if err != nil {
					return nil, err
				}

				slices = append(slices, names...)

				continue
			}
		}

		slices = append(slices, name)
	}

	return slices, nil
}

func getValuesFromValue(dataValue reflect.Value) (values []interface{}, err error) {

	slices := make([]interface{}, 0)

	dataType := dataValue.Type()
	for i := 0; i < dataValue.NumField(); i++ {

		field := dataValue.Field(i)
		if field.Type().Kind() == reflect.Struct {
			fieldType := dataType.Field(i)
			tagName := fieldType.Tag.Get(fieldTbl)
			if tagName == "" { // check if it has tag

				newSlices, err := getValuesFromValue(field) //recursive
				if err != nil {
					return nil, err
				}

				slices = append(slices, newSlices...)
				continue
			}

		}

		slices = append(slices, field.Interface())
	}

	return slices, nil
}

func inspectContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return nil
}

func createInsertCommand( /*ctx context.Context,*/ model IModel, markFunc func(int) (string, error)) (cmd string,
	values []interface{}, err error) {
	/*if err = inspectContext(ctx); err != nil {
		return
	}*/

	tblName, fields, values, err := extractFromModel(model)
	if err != nil {
		return
	}

	valuesMark, err := markFunc(len(fields))
	if err != nil {
		return
	}

	cmd = fmt.Sprintf("insert into %s(%s) values(%s)", tblName, strings.Join(fields, ","), valuesMark)
	return
}

func createUpdateCommand(model IModel, markFunc func(int) (string, error)) (cmd string,
	values []interface{}, err error) {

	tblName, fields, values, err := extractFromModel(model)
	if err != nil {
		return
	}

	//fmt.Println(fields, values)

	pairs := make([]string, 0)
	//var value interface{}
	//var ok bool

	for i, field := range fields {

		valueMark, err := markFunc(i)
		if err != nil {
			return "", nil, err
		}

		pairs = append(pairs, fmt.Sprintf("%s=%s", field, valueMark))
	}

	//fmt.Println(pairs)

	cmd = fmt.Sprintf("update %s set %s", tblName, strings.Join(pairs, ","))

	return cmd, values, nil
}

func createSelectCommand(model IModel) (cmd string, err error) {

	tblName, fields, _, err := extractFromModel(model)
	if err != nil {
		return
	}

	cmd = fmt.Sprintf("select %s from %s", strings.Join(fields, ","), tblName)

	return cmd, nil
}

func inspectResultOfSelect(result interface{}) (reflect.Type, error) {
	rstType := reflect.TypeOf(result)
	if rstType.Kind() != reflect.Ptr {
		return nil, errors.New("result must be pointer to slice of struct")
	}

	rstType = rstType.Elem()
	if rstType.Kind() != reflect.Slice {
		return nil, errors.New("result must be pointer to slice of struct")
	}

	if rstType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("result must be pointer to slice of struct")
	}

	return rstType, nil
}

func generateStorage(structData reflect.Value) []reflect.Value {
	storages := make([]reflect.Value, 0)

	structDataType := structData.Type()
	for i := 0; i < structData.NumField(); i++ {

		field := structData.Field(i)
		if field.Type().Kind() == reflect.Struct {

			fieldType := structDataType.Field(i)
			tagName := fieldType.Tag.Get(fieldTbl)
			if tagName == "" { // check if it has tag
				newStorages := generateStorage(field) //recursive
				storages = append(storages, newStorages...)
				continue
			}
		}

		storages = append(storages, field.Addr())
	}

	return storages
}
