package light

import "github.com/LuvDaSun/redux-go/redux"

/*
CreateApplicationStore does exactly that!
*/
func CreateApplicationStore() *redux.Store {
	reducer := func(state redux.State, action redux.Action) redux.State {
		return ReduceApplicationState(state.(*ApplicationState), action)
	}

	store := redux.CreateStore(InitialApplicationState, reducer)

	return store
}
