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
func CreateStore(reducer Reducer) *Store {
	state := reducer(nil, nil)

	dispatchHandler := func(store *Store, action Action) {
		store.state = reducer(store.state, action)
	}

	return &Store{
		state,
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
