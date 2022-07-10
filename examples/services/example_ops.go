package services

import (
	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
)

type exampleUsingOps struct {
	db   database.IConnection
	repo pvi.IPVIOperations
}

func NewExampleUsingOps(db database.IConnection, repo pvi.IPVIOperations) IExampleService {
	return &exampleUsingOps{db, repo}
}

func (svc *exampleUsingOps) GetAvailableEntities() ([]*pvi.PVIEntity, error) {
	tx, _ := svc.db.Begin()
	defer tx.Rollback()

	q := &pvi.Query{State: pvi.Available}
	results, err := svc.repo.Find(q)(tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return results, nil
}
