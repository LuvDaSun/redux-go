package light

import "github.com/LuvDaSun/redux-go/redux"

/*
ApplicationState main state
*/
type ApplicationState struct {
	isOn bool
}

/*
InitialApplicationState initial application state
*/
var InitialApplicationState = &ApplicationState{
	isOn: false,
}

/*
ReduceApplicationState will reduce application state
*/
func ReduceApplicationState(state *ApplicationState, action redux.Action) *ApplicationState {
	switch action.(type) {

	case SwitchOnAction:
		return &ApplicationState{
			isOn: true,
		}

	case SwitchOffAction:
		return &ApplicationState{
			isOn: false,
		}
	}

	return state
}
