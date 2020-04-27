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
Dispatch dispatches action
*/
type Dispatch func(Action)

/*
Store is a redux store
*/
type Store struct {
	DispatchChannel chan<- Action

	stateMutex    *sync.RWMutex
	state         State
	listenerMutex *sync.Mutex
	listenerIndex int
	listeners     map[int]Listener
	dispatch      Dispatch
}

/*
CreateStore creates a store
*/
func CreateStore(initalState State, reducer Reducer) *Store {
	dispatchChannel := make(chan Action)

	store := &Store{
		dispatchChannel,
		&sync.RWMutex{},
		initalState,
		&sync.Mutex{},
		0,
		make(map[int]Listener, 0),
		nil,
	}

	store.dispatch = func(action Action) {
		store.stateMutex.Lock()
		defer store.stateMutex.Unlock()

		store.state = reducer(store.state, action)

		store.listenerMutex.Lock()
		defer store.listenerMutex.Unlock()

		for _, listener := range store.listeners {
			listener(store.state)
		}

	}

	go store.dispatchLoop(dispatchChannel)

	return store
}

/*
GetState gets current of state
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

func (store *Store) dispatchLoop(dispatchChannel <-chan Action) {
	for action := range dispatchChannel {
		store.dispatch(action)
	}
}
