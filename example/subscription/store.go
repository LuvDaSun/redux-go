package subscription

import (
	"context"
	"time"

	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateApplicationStore does exactly that!
*/
func CreateApplicationStore(ctx context.Context, interval time.Duration) *redux.Store {
	reducer := func(state redux.State, action redux.Action) redux.State {
		return ReduceApplicationState(state.(*ApplicationState), action)
	}

	store := redux.CreateStore(InitialApplicationState, reducer).
		ApplyMiddleware(
			CreateSubscriptionMiddleware(ctx, interval),
		)

	return store
}
