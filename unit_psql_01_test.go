package databasex

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_Postgresql_CreateTable_01(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
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

func Test_Postgresql_AddOneRecord_02(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : add one record to tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) < 1 {
		t.Errorf("adding one data fail")
	}

	if data[0].ID != student.ID || data[0].Name != student.Name || data[0].Age != student.Age || data[0].Grade != student.Grade {
		t.Errorf("data is different")
	}

}

func Test_Postgresql_AddOneRecordWF_03(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : add one record to tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlFactory := CreateSQLFactory()
	sqlOp := sqlFactory.NewSQLOperation(postgres)

	t.Logf("delete all data first")

	mdlFactory := CreateModelFactory()
	model := mdlFactory.NewModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) < 1 {
		t.Errorf("adding one data fail")
	}

	if data[0].ID != student.ID || data[0].Name != student.Name || data[0].Age != student.Age || data[0].Grade != student.Grade {
		t.Errorf("data is different")
	}

}

func Test_Postgresql_UpdateOneRecord_04(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : update one records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName  = "eko win"
		CAge   = 8
		CGrade = 3
	)

	var keypair = struct {
		Name  string
		Age   int
		Grade int
	}{CName, CAge, CGrade}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateDb(context.Background(), model, "id='C001'"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "id='C001'", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if data[0].Name != CName || data[0].Age != CAge || data[0].Grade != CGrade {
		t.Errorf("data is different")
	}

}

func Test_Postgresql_DeleteOneRecord_05(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : delete one records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("delete one records from table")

	model.SetNewData(nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, "id='C001'"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data := make([]Student, 0)

	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "id='C001'", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) != 0 {
		t.Errorf("data is not deleted")
	}

}

func Test_Postgresql_UpdateSeveralRecords_06(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : update several records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName = "eko win"
		CAge  = 8
	)

	var keypair = struct {
		Name string
		Age  int
	}{CName, CAge}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateDb(context.Background(), model, "grade=5"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "grade=5", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) == 0 {
		t.Errorf("data is not inserted")
	}

	for _, val := range data {

		if val.Name != CName || val.Age != CAge {
			t.Errorf("data is different")
			break
		}

	}

}

func Test_Postgresql_UpdateAllRecords_07(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : update all records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName = "eko win"
		CAge  = 8
	)

	var keypair = struct {
		Name string
		Age  int
	}{CName, CAge}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) == 0 {
		t.Errorf("data is not inserted")
	}

	for _, val := range data {

		if val.Name != CName || val.Age != CAge {
			t.Errorf("data is different")
			break
		}

	}

}

func Test_Postgresql_DeleteAllRecords_08(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	t.Logf("testing : delete all records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 2}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("delete all records from table")

	model.SetNewData(nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) != 0 {
		t.Errorf("table still has data")
	}

}

func Test_Postgresql_DeleteAllRecords_09(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type AnakStudent struct {
		ID   string
		Name string
	}

	type Student struct {
		AnakStudent
		Age   int
		Grade int
	}

	t.Logf("testing : delete all records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{AnakStudent{"C001", "junjun"}, 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{}
	student.ID = "C002"
	student.Name = "maman"
	student.Age = 8
	student.Grade = 2

	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{AnakStudent{"C003", "yuli"}, 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	t.Logf("delete all records from table")

	model.SetNewData(nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read from table")

	data = make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) != 0 {
		t.Errorf("table still has data")
	}

}

func Test_Postgresql_DeleteAllRecords_10(t *testing.T) {

	log.Println(t.Name())

	t.Logf("testing : delete all records from tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_user", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	// TUser is type for user
	type TUser struct {
		Name     string `fieldtbl:"user_name"`
		Password string `fieldtbl:"password"`
		Email    string `fieldtbl:"email"`
		Role     string `fieldtbl:"role"`
	}

	// TUserTable is data type containing user data
	type TUserTable struct {
		TUser
		Status    string    `fieldtbl:"status"`
		CreatedAt time.Time `fieldtbl:"created_at"`
		//CreatedAt string `fieldtbl:"created_at"`
		LastUpdate time.Time `fieldtbl:"last_update"`
		//LastUpdate string `fieldtbl:"last_update"`
	}

	var userData TUserTable
	var userDataColl []TUserTable

	model = NewSimpleModel("tb_user", userData)

	err = sqlOp.SelectDb(context.Background(), model, "", &userDataColl)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(userDataColl) > 0 {
		t.Errorf("table should not have data any more\n")
	}

	t.Logf("add several data")

	username := "supriadi"
	password := "pass123"
	email := "supriadi@gmail.com"
	role := "umum"
	status := "pending"
	createdAt := time.Now()
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	lastUpdate := time.Now()
	//lastUpdate := time.Now().Format("2006-01-02 15:04:05")

	userData = TUserTable{TUser{username, password, email, role}, status, createdAt, lastUpdate}
	model.SetNewData(userData)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if err = sqlOp.SelectDb(context.Background(), model, "", &userDataColl); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if len(userDataColl) != 1 {
		t.Errorf("table should have only one record\n")
	}

	if err = sqlOp.SelectDb(context.Background(), model, "user_name='"+username+"'", &userDataColl); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if len(userDataColl) != 1 {
		t.Errorf("table should have only one record for user supriadi\n")
	}

	if userDataColl[0].Name != username ||
		userDataColl[0].Password != password ||
		userDataColl[0].Email != email ||
		userDataColl[0].Role != role ||
		userDataColl[0].Status != status ||
		userDataColl[0].CreatedAt.Unix() != createdAt.Unix() ||
		userDataColl[0].LastUpdate.Unix() != lastUpdate.Unix() {
		t.Errorf("data is different\n")
	}

	t.Logf("update status into 'active' and role into 'admin' ")

	username = "supriadi"
	//password := "pass123"
	//email := "supriadi@gmail.com"
	role = "admin"
	status = "active"
	//createdAt := time.Now()
	lastUpdate = time.Now()
	//lastUpdate = time.Now().Format("2006-01-02 15:04:05")

	type TUserUpdate struct {
		Role       string    `fieldtbl:"role"`
		Status     string    `fieldtbl:"status"`
		LastUpdate time.Time `fieldtbl:"last_update"`
		//LastUpdate string `fieldtbl:"last_update"`
	}

	updatedUser := TUserUpdate{role, status, lastUpdate}

	model.SetNewData(updatedUser)
	if _, err = sqlOp.UpdateDb(context.Background(), model, "user_name='"+username+"'"); err != nil {
		fmt.Println("update error : ", updatedUser)
		t.Errorf("%s\n", err.Error())
		return
	}

	t.Logf("Check update result\n")

	model.SetNewData(TUserTable{})
	if err = sqlOp.SelectDb(context.Background(), model, "user_name='"+username+"'", &userDataColl); err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if len(userDataColl) != 1 {
		t.Errorf("table should have only one record for user supriadi\n")
	}

	if userDataColl[0].Name != username ||
		userDataColl[0].Password != password ||
		userDataColl[0].Email != email ||
		userDataColl[0].Role != role ||
		userDataColl[0].Status != status ||
		userDataColl[0].CreatedAt.Unix() != createdAt.Unix() ||
		userDataColl[0].LastUpdate.Unix() != lastUpdate.Unix() {
		t.Errorf("data is different\n")
	}

	/*
		model.SetNewData(nil)
		if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		t.Logf("read from table")

		data = make([]Student, 0)
		model.SetNewData(Student{})
		if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		if len(data) != 0 {
			t.Errorf("table still has data")
		}
	*/
}

/*
func Test_Postgresql_datetime_11(t *testing.T) {

	log.Println(t.Name())

	// Student is type
	type Student struct {
		ID        string    `fieldtbl:"id"`
		Name      string    `fieldtbl:"name"`
		Age       int       `fieldtbl:"age"`
		Grade     int       `fieldtbl:"grade"`
		CreatedAt time.Time `fieldtbl:"created_at"`
	}

	t.Logf("testing : add date time field into tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1, time.Now()}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) < 1 {
		t.Errorf("adding one data fail")
	}

	if data[0].ID != student.ID || data[0].Name != student.Name || data[0].Age != student.Age ||
		data[0].Grade != student.Grade || isTimeDifferent(data[0].CreatedAt, student.CreatedAt) {
		t.Errorf("data is different")
	}

}*/

func Test_Postgresql_embed_12(t *testing.T) {

	log.Println(t.Name())

	type Kid struct {
		ID   string `fieldtbl:"id"`
		Name string `fieldtbl:"name"`
		Age  int    `fieldtbl:"age"`
	}

	// Student is type
	type Student struct {
		Kid
		Grade int `fieldtbl:"grade"`
	}

	t.Logf("testing : add date time field into tabel tb_student in db_belajar_golang database using postgresq")

	t.Logf("create connection to database server")

	postgres, err := NewPostgre(psqlUsernameTest, psqlPasswordTest, psqlHostTest, psqlPortTest, psqlDbTest,
		psqlOtherTest, psqlMaxConnectionsTest, psqlMaxIdleTest)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	sqlOp := NewSimpleSQL(postgres)

	t.Logf("delete all data first")

	model := NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{Kid{"C001", "junjun", 6}, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("read table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) < 1 {
		t.Errorf("adding one data fail")
	}

	if data[0].ID != student.ID || data[0].Name != student.Name || data[0].Age != student.Age || data[0].Grade != student.Grade {
		t.Errorf("data is different")
	}

}
