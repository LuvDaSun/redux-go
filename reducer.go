package redux

/*
Reducer will reduce a state and an action to a next state
*/
type Reducer func(State, Action) State
