package databasex

import (
	"fmt"
)

func ExampleNewSimpleModel() {

	type Student struct {
		Name  string
		Age   int
		Grade int
	}

	var student = Student{"Richard", 10, 5}

	// create model
	newModel := NewSimpleModel("tb_student", student)

	fmt.Println(newModel.GetTableName()) // output ==> tb_student
	fmt.Println(newModel.GetData())      // output ==> { Richard, 10, 5}

	student.Name = "Abraham"
	student.Age = 11
	student.Grade = 6

	// replace old data with new one
	newModel.SetNewData(student)

	fmt.Println(newModel.GetData()) // output ==> { Abraham, 11, 6}

	//Output:
	//tb_student
	//{Richard 10 5}
	//{Abraham 11 6}
}
