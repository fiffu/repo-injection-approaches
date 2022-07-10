package repoimpl

import (
	"errors"

	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
)

type stub struct {
	rows []*pvi.PVIEntity
}

func NewStubImpl(conn database.IPersistence) pvi.IPVIRepo {
	rows := []*pvi.PVIEntity{
		{
			ID:    "asd",
			State: pvi.Available,
		},
	}
	return &stub{rows}
}

func cmp(e *pvi.PVIEntity, q *pvi.Query) bool {
	if q.ID != "" && q.ID != e.ID {
		return false
	}
	if q.State != pvi.Unknown && q.State != e.State {
		return false
	}
	return true
}

func (s *stub) Find(q *pvi.Query) ([]*pvi.PVIEntity, error) {
	results := make([]*pvi.PVIEntity, 0)
	for _, e := range s.rows {
		if cmp(e, q) {
			results = append(results, e)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("no records found")
	}
	return results, nil
}

func (s *stub) FindOne(q *pvi.Query) (*pvi.PVIEntity, error) {
	found, err := s.Find(q)
	if err != nil {
		return nil, err
	}
	return found[0], nil
}

func (s *stub) Update(*pvi.PVIEntity, *pvi.Query, *pvi.UpdateValues) error {
	return nil
}
