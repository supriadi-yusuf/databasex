package databasex

//IModel is interface related to table. This interface has methods :
//
// - GetTableName() (tableName string)
//
//   This method is to get the table name associated to this model
//
// - GetData() (data interface{})
//
//   This method is to get data stored in this model.
//
type IModel interface {
	GetTableName() (tableName string)
	GetData() (data interface{})
}

// IWriteableModel is extension for IModel. It has new method SetNewData.
//
// Method SetNewData() is for replacing old data in the model with new one.
// This method has one input parameter. Value stored into this parameter must has struct type.
type IWriteableModel interface {
	IModel
	SetNewData(data interface{})
}

type simpleModel struct {
	name string
	data interface{}
}

func (t *simpleModel) GetTableName() string {
	return t.name
}

func (t *simpleModel) GetData() interface{} {
	return t.data
}

func (t *simpleModel) SetNewData(data interface{}) {
	t.data = data
}

// NewSimpleModel is function returning object whose type is IWriteableModel (interface).
// We need this object to work with table.
//
// This function receives two input parameter :
//
// - name
//   name of table to associted with model.
//
// - data
//   data is data we want to use in CRUD operation.
//   Value stored into data must has struct type.
func NewSimpleModel(name string, data interface{}) (model IWriteableModel) {
	var s simpleModel

	s.name = name
	s.data = data

	return &s
}
