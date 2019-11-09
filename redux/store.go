package redux

import (
	"sync"
)

/*
Store is a redux store
*/
type Store struct {
	state           State
	dispatchHandler DispatchHandler
	stateMutex      *sync.Mutex
	dispatchMutex   *sync.Mutex
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
		&sync.Mutex{},
	}
}

/*
Dispatch dispatches action
*/
func (store *Store) Dispatch(action Action) {
	store.dispatchMutex.Lock()
	defer store.dispatchMutex.Unlock()

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
