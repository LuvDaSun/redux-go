package redux

/*
Store is a redux store
*/
type Store struct {
	state           State
	dispatchHandler DispatchHandler
}

/*
CreateStore creates a store
*/
func CreateStore(initalState State, reducer Reducer) *Store {
	dispatchHandler := func(store *Store, action Action) {
		store.state = reducer(store.state, action)
	}

	return &Store{
		initalState,
		dispatchHandler,
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
	return store.state
}
