package subscription

import "github.com/LuvDaSun/redux-go/redux"

/*
ApplicationState state
*/
type ApplicationState struct {
	counter int
}

/*
InitialApplicationState initial application state
*/
var InitialApplicationState = &ApplicationState{
	counter: 0,
}

/*
ReduceApplicationState reduces application state
*/
func ReduceApplicationState(state *ApplicationState, action redux.Action) *ApplicationState {
	switch action.(type) {

	case *EventAction:
		return &ApplicationState{
			counter: state.counter + 1,
		}
	}

	return state
}
