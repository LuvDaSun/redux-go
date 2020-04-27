package redux

/*
StoreInterface interface to a store for middleware
*/
type StoreInterface struct {
	DispatchChannel chan<- Action

	store *Store
}

/*
Dispatch dispatches action
*/
func (store *StoreInterface) Dispatch(action Action) {
	store.store.dispatch(action)
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
type Middleware func(Dispatch) Dispatch

/*
ApplyMiddleware applies middleware to a store
*/
func (store *Store) ApplyMiddleware(middlewareFactories ...MiddlewareFactory) *Store {
	middlewares := make([]Middleware, len(middlewareFactories))
	for index, middlewareFactory := range middlewareFactories {
		middleware := middlewareFactory(StoreInterface{
			DispatchChannel: store.DispatchChannel,
			store:           store,
		})
		middlewares[index] = middleware
	}

	dispatchNext := store.dispatch
	for _, middleware := range middlewares {
		dispatchNext = middleware(dispatchNext)
	}
	store.dispatch = dispatchNext

	return store
}
