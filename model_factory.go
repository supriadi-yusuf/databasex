package databasex

// IModelFactory is interface defining how to create model.
type IModelFactory interface {
	NewModel(tableName string, data interface{}) (model IWriteableModel)
}

type modelFactory struct{}

func (mdl *modelFactory) NewModel(tableName string, data interface{}) (model IWriteableModel) {
	return NewSimpleModel(tableName, data)
}

// CreateModelFactory is function to create model factory
func CreateModelFactory() (factory IModelFactory) {
	return new(modelFactory)
}
