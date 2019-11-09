package redux

import (
	"sync"
)

/*
Action is a value that is dispatched to the redux store
*/
type Action interface{}

/*
Reducer will reduce a state and an action to a next state
*/
type Reducer func(State, Action) State

/*
State is the state of a redux store
*/
type State interface{}

/*
DispatchHandler will handle every dispatch
*/
type DispatchHandler func(
	store *Store,
	action Action,
)

/*
Store is a redux store
*/
type Store struct {
	state           State
	dispatchHandler DispatchHandler
	stateMutex      *sync.Mutex
}

/*
CreateStore creates a store
*/
func CreateStore(initalState State, reducer Reducer) *Store {
	dispatchHandler := func(store *Store, action Action) {
		store.stateMutex.Lock()
		defer store.stateMutex.Unlock()

		store.state = reducer(store.state, action)
	}

	return &Store{
		initalState,
		dispatchHandler,
		&sync.Mutex{},
	}
}

/*
Dispatch dispatches action
*/
func (store *Store) Dispatch(action Action) {
	store.dispatchHandler(store, action)
}

/*
GetState gets snapshot of state
*/
func (store *Store) GetState() State {
	store.stateMutex.Lock()
	defer store.stateMutex.Unlock()

	return store.state
}
