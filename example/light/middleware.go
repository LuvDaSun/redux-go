package light

import (
	"sync"

	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateToggleMiddleware transforms the toggle action
*/
func CreateToggleMiddleware() redux.MiddlewareFactory {

	return func(api redux.StoreInterface) redux.Middleware {
		return func(next redux.Dispatch) redux.Dispatch {
			var mutex = &sync.Mutex{}

			return func(action redux.Action) {
				next(action)

				switch action.(type) {
				case *ToggleAction:
					mutex.Lock()
					defer mutex.Unlock()

					state := api.GetState().(*ApplicationState)
					if state.SelectLightIsOn() {
						api.Dispatch(&SwitchOffAction{})
					} else {
						api.Dispatch(&SwitchOnAction{})
					}
				}

			}
		}
	}
}
