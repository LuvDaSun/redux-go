package light

import (
	"sync"
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

	store.Dispatch(&SwitchOnAction{})
	state2 := store.GetState().(*ApplicationState)

	store.Dispatch(&SwitchOffAction{})
	state3 := store.GetState().(*ApplicationState)

	store.Dispatch(&ToggleAction{})
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, false, state1.SelectLightIsOn())
	assert.Equal(test, true, state2.SelectLightIsOn())
	assert.Equal(test, false, state3.SelectLightIsOn())
	assert.Equal(test, true, state4.SelectLightIsOn())

	for range [100]int{} {
		var wg sync.WaitGroup

		job := func() {
			for range [100]int{} {
				store.Dispatch(&ToggleAction{})
			}
			wg.Done()
		}

		wg.Add(100)
		for range [100]int{} {
			go job()
		}

		wg.Wait()

		state5 := store.GetState().(*ApplicationState)
		assert.Equal(test, true, state5.SelectLightIsOn())
	}
}
