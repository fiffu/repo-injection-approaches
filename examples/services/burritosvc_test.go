package services

import (
	"testing"

	"github.com/fiffu/repo-injection-approaches/examples/domain/burrito"
	"github.com/fiffu/repo-injection-approaches/examples/infra/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetAvailableEntities(t *testing.T) {
	mockOps := &burrito.MockIBurritoOperations{}
	mockOps.On("Find", mock.Anything).Return(burrito.GetManyOp(func(conn database.IPersistence) ([]*burrito.BurritoEntity, error) {
		return []*burrito.BurritoEntity{
			{
				ID: "asd",
			},
		}, nil
	}))

	svc := NewBurritoService(database.NewDatabaseMock(), mockOps)
	results, err := svc.GetAvailable()

	assert.NoError(t, err)
	assert.Equal(t, "asd", results[0].ID)
}
