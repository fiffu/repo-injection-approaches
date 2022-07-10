package database

import "github.com/stretchr/testify/mock"

func NewDatabaseMock() *MockIConnection {
	tx := &MockITransaction{}
	tx.On("Select", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	tx.On("Exec", mock.Anything, mock.Anything).Return(nil)
	tx.On("Commit").Return(nil)
	tx.On("Rollback").Return(nil)

	conn := &MockIConnection{}
	conn.On("Begin").Return(tx, nil)
	return conn
}
