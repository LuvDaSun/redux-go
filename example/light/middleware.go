package light

import (
	"sync"

	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateToggleMiddleware transforms the toggle action
*/
func CreateToggleMiddleware() redux.MiddlewareFactory {

	return func(store redux.StoreInterface) redux.Middleware {
		return func(next redux.Dispatch) redux.Dispatch {
			var mutex = &sync.Mutex{}

			return func(action redux.Action) {
				next(action)

				switch action.(type) {
				case *ToggleAction:
					mutex.Lock()
					defer mutex.Unlock()

					state := store.GetState().(*ApplicationState)
					if state.SelectLightIsOn() {
						store.Dispatch(&SwitchOffAction{})
					} else {
						store.Dispatch(&SwitchOnAction{})
					}
				}

			}
		}
	}
}
