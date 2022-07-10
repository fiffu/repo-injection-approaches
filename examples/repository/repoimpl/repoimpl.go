package repoimpl

import (
	"strings"

	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
	"github.com/fiffu/repo-injection-approaches/examples/repository"
)

type impl struct {
	conn database.IPersistence
}

func NewRepositoryImpl(conn database.IPersistence) pvi.IPVIRepo {
	return &impl{conn}
}

func (i *impl) Find(q *pvi.Query) ([]*pvi.PVIEntity, error) {
	sql := "SELECT id, state FROM sometable {{whereFilters}}"

	// filter
	whereFilters := make([]string, 0)
	params := make([]interface{}, 0)
	if q.ID != "" {
		whereFilters = append(whereFilters, "id = ?")
		params = append(params, []byte(q.ID))
	}
	if q.State != pvi.Unknown {
		whereFilters = append(whereFilters, "state = ?")
		params = append(params, repository.EncodeEntityState(q.State))
	}
	if len(whereFilters) > 0 {
		sql = strings.ReplaceAll(sql, "{{whereFilters}}", strings.Join(whereFilters, " AND "))
	}

	// query
	rows := make([]*repository.EntityDBModel, 0)
	err := i.conn.Select(&rows, sql, params...)
	if err != nil {
		return nil, err
	}

	// deserialize
	result := make([]*pvi.PVIEntity, 0)
	for _, row := range rows {
		result = append(result, repository.Deserialize(row))
	}

	return result, nil
}

func (i *impl) FindOne(q *pvi.Query) (*pvi.PVIEntity, error) {
	found, err := i.Find(q)
	if err != nil {
		return nil, err
	}
	return found[0], nil
}

func (i *impl) Update(*pvi.PVIEntity, *pvi.Query, *pvi.UpdateValues) error {
	return nil
}
