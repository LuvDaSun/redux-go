package redux

/*
MiddlewareFactory is the middleware factory :-)
*/
type MiddlewareFactory func(GetState, Dispatch) Middleware

/*
GetState gets state
*/
type GetState func() State

/*
Dispatch dispatches action
*/
type Dispatch func(Action)

/*
Middleware is the actual middleware
*/
type Middleware func(Dispatch) Dispatch

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewareFactories ...MiddlewareFactory) *Store {
	dispatchHandler := store.dispatchHandler

	getState := func() State {
		return store.GetState()
	}

	dispatch := func(action Action) {
		dispatchHandler(store, action)
	}

	middlewares := make([]Middleware, len(middlewareFactories))
	for index, middlewareFactory := range middlewareFactories {
		middleware := middlewareFactory(getState, dispatch)
		middlewares[index] = middleware
	}

	dispatchNext := dispatch
	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}

	nextDispatchHandler := func(store *Store, action Action) {
		dispatchNext(action)
	}

	store.dispatchHandler = nextDispatchHandler

	return store
}
