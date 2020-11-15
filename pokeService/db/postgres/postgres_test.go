package postgres

import (
	"database/sql"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/Enviy/pokedexProject/pokeService/entity"
)

var (
	defaultSecrets = entity.Secrets{
		PostgresInfo: entity.PostgresInfo{
			Name:     "test",
			User:     "test user",
			Password: "",
			Port:     1234,
			Host:     "localhost",
		},
	}
	errFoo = errors.New("errored")
)

func setupFakeOpenDBConn(respErr error) func() {
	original := openDBConnection
	openDBConnection = func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, respErr
	}

	return func() {
		openDBConnection = original
	}
}

func Test_New(t *testing.T) {
	tests := []struct {
		Name        string
		Secrets     entity.Secrets
		OpenConnErr error

		Err error
	}{
		{
			Name:    "Success",
			Secrets: defaultSecrets,
		},
		{
			Name:        "Error",
			Secrets:     defaultSecrets,
			OpenConnErr: errFoo,

			Err: errFoo,
		},
	}

	for i := 0; i < len(tests); i++ {
		currTest := tests[i]
		t.Run(currTest.Name, func(t *testing.T) {
			cleanup := setupFakeOpenDBConn(currTest.OpenConnErr)
			defer cleanup()

			pg, err := New(currTest.Secrets)
			if currTest.Err != nil {
				assert.Error(t, err)
				assert.Nil(t, pg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, pg)
			}
		})
	}
}
