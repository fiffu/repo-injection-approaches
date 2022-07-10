package pvi

import "github.com/fiffu/repo-injection-approaches/examples/database"

type IPVIRepo interface {
	Find(*Query) ([]*PVIEntity, error)
	FindOne(*Query) (*PVIEntity, error)
	Update(*PVIEntity, *Query, *UpdateValues) error
}

type GetManyOp func(conn database.IPersistence) ([]*PVIEntity, error)
type GetOneOp func(conn database.IPersistence) (*PVIEntity, error)
type UpdateOp func(conn database.IPersistence) error

type IPVIOperations interface {
	Find(*Query) GetManyOp
	FindOne(*Query) GetOneOp
	Update(*PVIEntity, *Query, *UpdateValues) UpdateOp
}
