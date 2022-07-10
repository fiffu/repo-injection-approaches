package database

type IPersistence interface {
	Select(dest interface{}, sqlQuery string, params ...interface{}) error
	Exec(sqlQuery string, params ...interface{}) error
}

type IConnection interface {
	IPersistence

	Begin() (ITransaction, error)
}

type ITransaction interface {
	IPersistence

	Commit() error
	Rollback() error
}
