package subscription

import (
	"context"
	"time"

	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateSubscriptionMiddleware emits events!
*/
func CreateSubscriptionMiddleware(ctx context.Context, interval time.Duration) redux.MiddlewareFactory {

	return func(store redux.StoreInterface) redux.Middleware {
		cancels := make(map[int]context.CancelFunc)

		dispatchEventAction := func(subscriptionCtx context.Context, id int) {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()

			for {
				select {
				case <-subscriptionCtx.Done():
					return

				case <-ticker.C:
					store.Dispatch(&EventAction{
						ID: id,
					})
				}
			}
		}
		handleStartAction := func(id int) {
			subscriptionCtx, cancel := context.WithCancel(ctx)
			cancels[id] = cancel
			go dispatchEventAction(subscriptionCtx, id)

		}
		handleStopAction := func(id int) {
			cancel := cancels[id]
			delete(cancels, id)
			cancel()
		}

		return func(next redux.Dispatcher) redux.Dispatcher {

			return func(action redux.Action) {
				switch action := action.(type) {
				case *StartAction:
					handleStartAction(action.ID)
					next(action)

				case *StopAction:
					next(action)
					handleStopAction(action.ID)

				default:
					next(action)
				}

			}
		}
	}
}
