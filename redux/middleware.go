package redux

/*
Middleware is the middleware API :-)
*/
type Middleware func(GetState, Dispatch) Chain

/*
GetState gets state
*/
type GetState func() State

/*
Dispatch dispatches action
*/
type Dispatch func(Action)

/*
Chain dispatches action
*/
type Chain func(Dispatch) Dispatch

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewares ...Middleware) *Store {
	dispatchHandler := store.dispatchHandler

	getState := func() State {
		return store.GetState()
	}

	dispatch := func(action Action) {
		dispatchHandler(store, action)
	}

	chains := make([]Chain, len(middlewares))
	for index, middleware := range middlewares {
		chain := middleware(getState, dispatch)
		chains[index] = chain
	}

	dispatchNext := dispatch
	for _, chain := range chains {
		dispatchNext = chain(dispatchNext)
	}

	nextDispatchHandler := func(store *Store, action Action) {
		dispatchNext(action)
	}

	store.dispatchHandler = nextDispatchHandler

	return store
}
