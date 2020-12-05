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

		if field.Type.Kind() == reflect.Struct {
			names, err := getFieldsFromType(field.Type)
			if err != nil {
				return nil, err
			}

			slices = append(slices, names...)
			continue
		}

		name := field.Name
		tagName := field.Tag.Get(fieldTbl)
		if tagName != "" {
			name = tagName
		}

		slices = append(slices, name)
	}

	return slices, nil
}

func getValuesFromValue(dataValue reflect.Value) (values []interface{}, err error) {

	slices := make([]interface{}, 0)

	for i := 0; i < dataValue.NumField(); i++ {

		field := dataValue.Field(i)
		if field.Type().Kind() == reflect.Struct {
			newSlices, err := getValuesFromValue(field)
			if err != nil {
				return nil, err
			}

			slices = append(slices, newSlices...)
			continue
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

func createInsertCommand( /*ctx context.Context,*/ model IModel, markFunc func(int) (string, error)) (cmd string, values []interface{}, err error) {
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

func createUpdateCommand(model IModel) (cmd string, err error) {

	tblName, fields, values, err := extractFromModel(model)
	if err != nil {
		return
	}

	//fmt.Println(fields, values)

	pairs := make([]string, 0)
	var value interface{}
	var ok bool

	for idx, field := range fields {
		value = values[idx]
		_, ok = value.(string)
		if ok {
			value = fmt.Sprintf("'%v'", value)
		}

		//fmt.Printf("%s=%v\n", field, value)
		//fmt.Println(value)

		pairs = append(pairs, fmt.Sprintf("%s=%v", field, value))

	}

	//fmt.Println(pairs)

	cmd = fmt.Sprintf("update %s set %s", tblName, strings.Join(pairs, ","))

	return cmd, nil
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
