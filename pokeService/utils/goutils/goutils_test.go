package goutils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrWrap(t *testing.T) {
	err := ErrWrap(errors.New("errored"))
	assert.Equal(t, "testing.tRunner: errored", err.Error())
}
