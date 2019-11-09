package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	store := CreateApplicationStore()

	state1 := store.GetState().(*ApplicationState)

	assert.NotNil(test, state1)
}
