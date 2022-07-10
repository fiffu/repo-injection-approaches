package pvi

import "github.com/fiffu/repo-injection-approaches/examples/database"

//go:generate mockery --name IPVIOperations --inpackage

type GetManyOp func(conn database.IPersistence) ([]*PVIEntity, error)
type GetOneOp func(conn database.IPersistence) (*PVIEntity, error)
type UpdateOp func(conn database.IPersistence) error

type IPVIOperations interface {
	Find(*Query) GetManyOp
	FindOne(*Query) GetOneOp
	Update(*PVIEntity, *Query, *UpdateValues) UpdateOp
}
