package databasex

// ISqlFactory is interface for creating factory of object implementing ISqlOperation
type ISqlFactory interface {
	NewSQLOperation(db IDatabase) ISqlOperation
}

type sqlFactory struct{}

func (factory *sqlFactory) NewSQLOperation(db IDatabase) ISqlOperation {
	return NewSimpleSQL(db)
}

// CreateSQLFactory is function to create factory of object implementing ISqlOperation
func CreateSQLFactory() ISqlFactory {
	return new(sqlFactory)
}
