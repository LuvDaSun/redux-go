package redux

/*
DispatchHandler will handle every dispatch
*/
type DispatchHandler func(
	store *Store,
	action Action,
)
