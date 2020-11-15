package secrets

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

var (
	testingYAMLPath = "./secrets.yaml"
	errFoo          = errors.New("errored")
)

func fakeYamlPathForTest(path string) func() {
	original := secretsYAMLPath
	secretsYAMLPath = path

	return func() {
		secretsYAMLPath = original
	}
}

func Test_New(t *testing.T) {

	badYamlFile, err := ioutil.TempFile("", "")
	badYamlFile.WriteString("postgres:\n  port: \"5432\"")
	assert.NoError(t, err)
	defer os.Remove(badYamlFile.Name())

	tests := []struct {
		Name     string
		FilePath string

		Err error
	}{
		{
			Name:     "Success",
			FilePath: testingYAMLPath,
		},
		{
			Name:     "Error from file",
			FilePath: "./bad.yaml",
			Err:      errFoo,
		},
		{
			Name:     "Error populating entity",
			FilePath: badYamlFile.Name(),
			Err:      errFoo,
		},
	}

	for i := 0; i < len(tests); i++ {
		currTest := tests[i]
		t.Run(currTest.Name, func(t *testing.T) {
			cleanup := fakeYamlPathForTest(currTest.FilePath)
			defer cleanup()

			secrets, err := New()
			if currTest.Err != nil {
				assert.Error(t, err)
				assert.Zero(t, secrets)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, secrets)
			}
		})
	}
}
