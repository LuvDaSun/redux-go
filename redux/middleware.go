package redux

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
		middleware := middlewareFactory(store)
		middlewares[index] = middleware
	}

	dispatchNext := store.dispatcher
	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}
	store.dispatcher = dispatchNext

	return store
}
