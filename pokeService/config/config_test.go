package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"

	"github.com/Enviy/pokedexProject/pokeService/constants"

	"github.com/stretchr/testify/assert"
)

var (
	errFoo = errors.New("errored")
)

func setupTestYAMLPaths(baseP, prodP string) func() {
	originalBase := baseYAMLPath
	baseYAMLPath = baseP
	originalProd := productionYAMLPath
	productionYAMLPath = prodP

	return func() {
		baseYAMLPath = originalBase
		productionYAMLPath = originalProd
	}
}

func setupServerEnv(env string) func() {
	original := serverEnvVariable
	serverEnvVariable = "TESTING_ENV_VARIABLE"
	originalValue := os.Getenv(serverEnvVariable)
	os.Setenv(serverEnvVariable, env)

	return func() {
		os.Setenv(serverEnvVariable, originalValue)
		serverEnvVariable = original
	}
}

func Test_New(t *testing.T) {
	defaultBasePath := "./base.yaml"
	defaultProdPath := "./production.yaml"

	badYamlFile, err := ioutil.TempFile("", "")
	badYamlFile.WriteString("serverPort: \"1234\"")
	assert.NoError(t, err)
	defer os.Remove(badYamlFile.Name())

	tests := []struct {
		Name      string
		BasePath  string
		ProdPath  string
		ServerEnv string

		Err error
	}{
		{
			Name:      "Success in development",
			BasePath:  defaultBasePath,
			ProdPath:  defaultProdPath,
			ServerEnv: constants.EnvDevelopment,
		},
		{
			Name:      "Success in production",
			BasePath:  defaultBasePath,
			ProdPath:  defaultProdPath,
			ServerEnv: constants.EnvProduction,
		},
		{
			Name:      "Error in development - bad path",
			BasePath:  "./badpath.yaml",
			ProdPath:  defaultProdPath,
			ServerEnv: constants.EnvDevelopment,

			Err: errFoo,
		},
		{
			Name:      "Error invalid yaml file",
			BasePath:  badYamlFile.Name(),
			ProdPath:  defaultProdPath,
			ServerEnv: constants.EnvDevelopment,

			Err: errFoo,
		},
	}

	for i := 0; i < len(tests); i++ {
		currTest := tests[i]
		t.Run(currTest.Name, func(t *testing.T) {
			cleanup := setupTestYAMLPaths(currTest.BasePath, currTest.ProdPath)
			cleanupEnv := setupServerEnv(currTest.ServerEnv)
			defer cleanup()
			defer cleanupEnv()

			config, err := New()
			if currTest.Err != nil {
				assert.Error(t, err)
				assert.Zero(t, config)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, config)
				assert.Equal(t, currTest.ServerEnv, config.Env)
			}
		})
	}
}
