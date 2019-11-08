package redux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ApplicationState struct {
	counter int
}

func (state ApplicationState) selectCounter() int {
	return state.counter
}

type IncrementAction struct{}
type DecrementAction struct{}

func ReduceApplicationState(state State, action Action) State {
	if state == nil {
		return ApplicationState{
			counter: 0,
		}
	}

	if action == nil {
		return state
	}

	{
		s := state.(ApplicationState)
		switch action.(type) {

		case IncrementAction:
			return ApplicationState{
				counter: s.counter + 1,
			}

		case DecrementAction:
			return ApplicationState{
				counter: s.counter - 1,
			}
		}
	}

	return state
}

func TestStore(test *testing.T) {
	store := CreateStore(ReduceApplicationState)

	state1 := store.GetState()

	store.Dispatch(IncrementAction{})

	state2 := store.GetState()

	store.Dispatch(DecrementAction{})

	state3 := store.GetState()

	assert.NotEqual(test, state1, state2)
	assert.NotEqual(test, state2, state3)
	assert.Equal(test, state3, state1)
}
