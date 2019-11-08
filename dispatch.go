package redux

/*
DispatchHandler will handle every dispatch
*/
type DispatchHandler func(
	store *Store,
	action Action,
)

func defaultDispatchHandler(store *Store, action Action) {
	store.state = store.reducer(store.state, action)
}
