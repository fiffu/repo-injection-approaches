package services

import (
	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
	"github.com/fiffu/repo-injection-approaches/examples/repository/repoimpl"
)

type IExampleService interface {
	GetAvailableEntities() ([]*pvi.PVIEntity, error)
}

type exampleService struct {
	db database.IConnection
}

func NewExample(conn database.IConnection) IExampleService {
	return &exampleService{conn}
}

func (svc *exampleService) GetAvailableEntities() ([]*pvi.PVIEntity, error) {
	tx, _ := svc.db.Begin()
	defer tx.Rollback()

	repo := repoimpl.NewRepositoryImpl(tx)

	q := &pvi.Query{State: pvi.Available}
	results, err := repo.Find(q)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return results, nil
}
