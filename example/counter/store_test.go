package counter

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

	store := redux.CreateStore(InitialApplicationState, reducer)

	state1 := store.GetState().(*ApplicationState)

	store.Dispatch(&IncrementAction{})
	state2 := store.GetState().(*ApplicationState)

	store.Dispatch(&IncrementAction{})
	state3 := store.GetState().(*ApplicationState)

	store.Dispatch(&DecrementAction{})
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, 0, state1.SelectCounter())
	assert.Equal(test, 1, state2.SelectCounter())
	assert.Equal(test, 2, state3.SelectCounter())
	assert.Equal(test, 1, state4.SelectCounter())

	var wg sync.WaitGroup

	job := func() {
		for range [1000]int{} {
			store.Dispatch(&IncrementAction{})
		}
		wg.Done()
	}

	wg.Add(1000)
	for range [1000]int{} {
		go job()
	}

	wg.Wait()

	state5 := store.GetState().(*ApplicationState)
	assert.Equal(test, 1+1000*1000, state5.SelectCounter())
}
