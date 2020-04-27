package redux

/*
StoreInterface interface to a store for middleware
*/
type StoreInterface struct {
	store *Store
}

/*
Dispatch dispatches action
*/
func (store *StoreInterface) Dispatch(action Action) {
	store.store.Dispatch(action)
}

/*
GetState gets current of state
*/
func (store *StoreInterface) GetState() State {
	return store.store.GetState()
}

/*
MiddlewareFactory is the middleware factory :-)
*/
type MiddlewareFactory func(StoreInterface) Middleware

/*
Middleware is the actual middleware
*/
type Middleware func(Dispatcher) Dispatcher

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewareFactories ...MiddlewareFactory) *Store {
	middlewares := make([]Middleware, len(middlewareFactories))
	for index, middlewareFactory := range middlewareFactories {
		middleware := middlewareFactory(StoreInterface{
			store: store,
		})
		middlewares[index] = middleware
	}

	dispatchNext := store.dispatcher
	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}
	store.dispatcher = dispatchNext

	return store
}
