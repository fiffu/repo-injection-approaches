package services

import (
	"testing"

	"github.com/fiffu/repo-injection-approaches/examples/database"
	"github.com/fiffu/repo-injection-approaches/examples/domain/pvi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetAvailableEntities(t *testing.T) {
	mockOps := &pvi.MockIPVIOperations{}
	mockOps.On("Find", mock.Anything).Return(pvi.GetManyOp(func(conn database.IPersistence) ([]*pvi.PVIEntity, error) {
		return []*pvi.PVIEntity{
			{
				ID: "asd",
			},
		}, nil
	}))

	svc := NewExampleUsingOps(database.NewDatabaseMock(), mockOps)
	results, err := svc.GetAvailableEntities()

	assert.NoError(t, err)
	assert.Equal(t, "asd", results[0].ID)
}
