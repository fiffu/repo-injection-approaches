package operationsimpl

import (
	"strings"

	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
	"github.com/fiffu/repo-injection-approaches/examples/repository"
)

type operationsImpl struct{}

func NewOperationsRepo() pvi.IPVIOperations {
	return &operationsImpl{}
}

func (i *operationsImpl) Find(q *pvi.Query) pvi.GetManyOp {
	return func(conn database.IPersistence) ([]*pvi.PVIEntity, error) {
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
		err := conn.Select(&rows, sql, params...)
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
}

func (i *operationsImpl) FindOne(q *pvi.Query) pvi.GetOneOp {
	return func(conn database.IPersistence) (*pvi.PVIEntity, error) {
		found, err := i.Find(q)(conn)
		if err != nil {
			return nil, err
		}
		return found[0], nil
	}
}

func (i *operationsImpl) Update(*pvi.PVIEntity, *pvi.Query, *pvi.UpdateValues) pvi.UpdateOp {
	return func(conn database.IPersistence) error {
		return nil
	}
}
