package light

import (
	"github.com/LuvDaSun/redux-go/redux"
)

/*
CreateToggleMiddleware transforms the toggle action
*/
func CreateToggleMiddleware() redux.MiddlewareFactory {

	return func(store redux.StoreInterface) redux.Middleware {
		return func(next redux.Dispatch) redux.Dispatch {
			return func(action redux.Action) {
				next(action)

				switch action.(type) {
				case *ToggleAction:
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
