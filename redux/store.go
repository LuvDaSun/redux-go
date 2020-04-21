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
	stateMutex    *sync.RWMutex
	state         State
	listenerMutex *sync.Mutex
	listenerIndex int
	listeners     map[int]Listener
	dispatcher    Dispatcher
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
		&sync.RWMutex{},
		initalState,
		&sync.Mutex{},
		0,
		make(map[int]Listener, 0),
		nil,
	}

	store.dispatcher = func(action Action) {
		store.stateMutex.Lock()
		defer store.stateMutex.Unlock()

		store.state = reducer(store.state, action)

		store.listenerMutex.Lock()
		defer store.listenerMutex.Unlock()

		for _, listener := range store.listeners {
			listener(store.state)
		}

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

/*
Listener function that listens to a store
*/
type Listener func(state State)

/*
Unsubscribe unsibscribed the listener
*/
type Unsubscribe func()

/*
Subscribe subscribes to the stores state
*/
func (store *Store) Subscribe(listener Listener) Unsubscribe {
	store.listenerMutex.Lock()
	defer store.listenerMutex.Unlock()

	store.listenerIndex++
	listenerIndex := store.listenerIndex

	store.listeners[listenerIndex] = listener
	unsubscribe := func() {
		store.listenerMutex.Lock()
		defer store.listenerMutex.Unlock()

		delete(store.listeners, listenerIndex)
	}

	return unsubscribe
}
