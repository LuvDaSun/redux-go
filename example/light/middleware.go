package light

import (
	"sync"

	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateToggleMiddleware transforms the toggle action
*/
func CreateToggleMiddleware() redux.Middleware {

	return func(getState redux.GetState, dispatch redux.Dispatch) redux.Chain {
		return func(next redux.Dispatch) redux.Dispatch {
			var mutex = &sync.Mutex{}

			return func(action redux.Action) {
				next(action)

				switch action.(type) {
				case *ToggleAction:
					mutex.Lock()
					defer mutex.Unlock()

					state := getState().(*ApplicationState)
					if state.SelectLightIsOn() {
						dispatch(&SwitchOffAction{})
					} else {
						dispatch(&SwitchOnAction{})
					}
				}

			}
		}
	}
}
