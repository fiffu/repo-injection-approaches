package services

import (
	"github.com/fiffu/repo-injection-approaches/examples/domain/burrito"
	"github.com/fiffu/repo-injection-approaches/examples/infra/database"
)

type IBurritoService interface {
	GetAvailable() ([]*burrito.BurritoEntity, error)
}

type burritoSvc struct {
	db   database.IConnection
	repo burrito.IBurritoOperations
}

func NewBurritoService(db database.IConnection, repo burrito.IBurritoOperations) IBurritoService {
	return &burritoSvc{db, repo}
}

func (svc *burritoSvc) GetAvailable() ([]*burrito.BurritoEntity, error) {
	tx, _ := svc.db.Begin()
	defer tx.Rollback()

	q := &burrito.Query{State: burrito.Available}
	results, err := svc.repo.Find(q)(tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return results, nil
}
