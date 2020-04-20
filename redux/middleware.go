package redux

/*
MiddlewareFactory is the middleware factory :-)
*/
type MiddlewareFactory func(StoreInterface) Middleware

/*
Middleware is the actual middleware
*/
type Middleware func(Dispatch) Dispatch

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewareFactories ...MiddlewareFactory) *Store {
	middlewares := make([]Middleware, len(middlewareFactories))
	for index, middlewareFactory := range middlewareFactories {
		middleware := middlewareFactory(store)
		middlewares[index] = middleware
	}

	dispatchNext := store.dispatchHandler
	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}
	store.dispatchHandler = dispatchNext

	return store
}
