package redux

/*
MiddlewareAPI exposes GetState and Dispatch of the store
*/
type MiddlewareAPI struct {
	store *Store
}

/*
MiddlewareFactory is the middleware factory :-)
*/
type MiddlewareFactory func(*MiddlewareAPI) Middleware

/*
GetState gets state
*/
type GetState func() State

/*
Middleware is the actual middleware
*/
type Middleware func(Dispatch) Dispatch

/*
GetState gets the state of the store
*/
func (middlewareAPI *MiddlewareAPI) GetState() State {
	return middlewareAPI.store.GetState()
}

/*
Dispatch dispatches action to the store
*/
func (middlewareAPI *MiddlewareAPI) Dispatch(action Action) {
	middlewareAPI.store.Dispatch(action)
}

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewareFactories ...MiddlewareFactory) *Store {
	middlewareAPI := &MiddlewareAPI{
		store: store,
	}

	middlewares := make([]Middleware, len(middlewareFactories))
	for index, middlewareFactory := range middlewareFactories {
		middleware := middlewareFactory(middlewareAPI)
		middlewares[index] = middleware
	}

	dispatchNext := store.dispatchHandler

	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}

	nextDispatchHandler := func(action Action) {
		dispatchNext(action)
	}
	store.dispatchHandler = nextDispatchHandler

	return store
}
