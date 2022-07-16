package burrito

import "github.com/fiffu/repo-injection-approaches/examples/infra/database"

//go:generate mockery --name IBurritoOperations --inpackage

type GetManyOp func(conn database.IPersistence) ([]*BurritoEntity, error)
type GetOneOp func(conn database.IPersistence) (*BurritoEntity, error)
type UpdateOp func(conn database.IPersistence) error

type IBurritoOperations interface {
	Find(*Query) GetManyOp
	FindOne(*Query) GetOneOp
	Update(*BurritoEntity, *Query, *UpdateValues) UpdateOp
}
