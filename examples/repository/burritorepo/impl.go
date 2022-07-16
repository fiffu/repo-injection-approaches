package burritorepo

import (
	"strings"

	"github.com/fiffu/repo-injection-approaches/examples/domain/burrito"
	"github.com/fiffu/repo-injection-approaches/examples/infra/database"
)

type operationsImpl struct{}

func NewOperationsRepo() burrito.IBurritoOperations {
	return &operationsImpl{}
}

func (i *operationsImpl) Find(q *burrito.Query) burrito.GetManyOp {
	return func(conn database.IPersistence) ([]*burrito.BurritoEntity, error) {
		sql := "SELECT id, state FROM sometable {{whereFilters}}"

		// filter
		whereFilters := make([]string, 0)
		params := make([]interface{}, 0)
		if q.ID != "" {
			whereFilters = append(whereFilters, "id = ?")
			params = append(params, []byte(q.ID))
		}
		if q.State != burrito.Unknown {
			whereFilters = append(whereFilters, "state = ?")
			params = append(params, EncodeEntityState(q.State))
		}
		if len(whereFilters) > 0 {
			sql = strings.ReplaceAll(sql, "{{whereFilters}}", strings.Join(whereFilters, " AND "))
		}

		// query
		rows := make([]*BurritoDBModel, 0)
		err := conn.Select(&rows, sql, params...)
		if err != nil {
			return nil, err
		}

		// deserialize
		result := make([]*burrito.BurritoEntity, 0)
		for _, row := range rows {
			result = append(result, Deserialize(row))
		}

		return result, nil
	}
}

func (i *operationsImpl) FindOne(q *burrito.Query) burrito.GetOneOp {
	return func(conn database.IPersistence) (*burrito.BurritoEntity, error) {
		found, err := i.Find(q)(conn)
		if err != nil {
			return nil, err
		}
		return found[0], nil
	}
}

func (i *operationsImpl) Update(*burrito.BurritoEntity, *burrito.Query, *burrito.UpdateValues) burrito.UpdateOp {
	return func(conn database.IPersistence) error {
		return nil
	}
}
