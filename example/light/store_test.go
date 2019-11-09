package light

import (
	"testing"

	"github.com/LuvDaSun/redux-go/redux"
	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	reducer := func(state redux.State, action redux.Action) redux.State {
		return ReduceApplicationState(state.(*ApplicationState), action)
	}

	store := redux.CreateStore(InitialApplicationState, reducer).
		ApplyMiddleware(CreateToggleMiddleware())

	state1 := store.GetState().(*ApplicationState)

	store.Dispatch(SwitchOnAction{})
	state2 := store.GetState().(*ApplicationState)

	store.Dispatch(SwitchOffAction{})
	state3 := store.GetState().(*ApplicationState)

	store.Dispatch(ToggleAction{})
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, false, state1.SelectLightIsOn())
	assert.Equal(test, true, state2.SelectLightIsOn())
	assert.Equal(test, false, state3.SelectLightIsOn())
	assert.Equal(test, true, state4.SelectLightIsOn())
}
