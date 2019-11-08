package redux

/*
Store is a redux store
*/
type Store struct {
	state           State
	reducer         Reducer
	dispatchHandler DispatchHandler
}

/*
CreateStore creates a store
*/
func CreateStore(reducer Reducer) Store {
	state := reducer(nil, nil)
	return Store{
		state,
		reducer,
		defaultDispatchHandler,
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
