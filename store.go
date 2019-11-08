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
	mutex           sync.Mutex
}

/*
CreateStore creates a store
*/
func CreateStore(reducer Reducer) *Store {
	state := reducer(nil, nil)

	dispatchHandler := func(store *Store, action Action) {
		store.mutex.Lock()
		defer store.mutex.Unlock()

		store.state = reducer(store.state, action)
	}

	return &Store{
		state,
		dispatchHandler,
		sync.Mutex{},
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
	store.mutex.Lock()
	defer store.mutex.Unlock()

	return store.state
}
