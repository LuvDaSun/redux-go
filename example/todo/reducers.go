package todo

import "github.com/LuvDaSun/redux-go/redux"

/*
ApplicationState state
*/
type ApplicationState struct {
}

/*
InitialApplicationState initial application state
*/
var InitialApplicationState = &ApplicationState{}

/*
ReduceApplicationState reduces application state
*/
func ReduceApplicationState(state *ApplicationState, action redux.Action) *ApplicationState {
	switch action.(type) {

	}

	return state
}
