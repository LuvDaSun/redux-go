package redux

import (
	"sync"
)

/*
State is the state of a redux store
*/
type State interface{}

/*
Action is a value that is dispatched to the redux store
*/
type Action interface{}

/*
Reducer will reduce a state and an action to a next state
*/
type Reducer func(State, Action) State

/*
Dispatcher dispatches action
*/
type Dispatcher func(Action)

/*
Store is a redux store
*/
type Store struct {
	state      State
	dispatcher Dispatcher
	stateMutex *sync.RWMutex
}

/*
StoreInterface defines methods
*/
type StoreInterface interface {
	GetState() State
	Dispatch(Action)
}

/*
CreateStore creates a store
*/
func CreateStore(initalState State, reducer Reducer) *Store {
	store := &Store{
		initalState,
		nil,
		&sync.RWMutex{},
	}

	store.dispatcher = func(action Action) {
		store.stateMutex.Lock()
		defer store.stateMutex.Unlock()

		store.state = reducer(store.state, action)
	}

	return store
}

/*
Dispatch dispatches action
*/
func (store *Store) Dispatch(action Action) {
	store.dispatcher(action)
}

/*
GetState gets snapshot of state
*/
func (store *Store) GetState() State {
	store.stateMutex.RLock()
	defer store.stateMutex.RUnlock()

	return store.state
}
