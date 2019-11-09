package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	store := CreateApplicationStore()

	state1 := store.GetState().(*ApplicationState)

	store.Dispatch(&AddTaskItemAction{
		id:   "a",
		name: "do a thing",
	})
	state2 := store.GetState().(*ApplicationState)

	store.Dispatch(&CompleteTaskItemAction{
		id: "a",
	})
	state3 := store.GetState().(*ApplicationState)

	store.Dispatch(&RemoveTaskItemAction{
		id: "a",
	})
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, 0, state1.SelectTaskCount())
	assert.Equal(test, 0, state1.SelectTaskCompleteCount())

	assert.Equal(test, 1, state2.SelectTaskCount())
	assert.Equal(test, 0, state2.SelectTaskCompleteCount())

	assert.Equal(test, 1, state3.SelectTaskCount())
	assert.Equal(test, 1, state3.SelectTaskCompleteCount())

	assert.Equal(test, 0, state4.SelectTaskCount())
	assert.Equal(test, 0, state4.SelectTaskCompleteCount())
}
