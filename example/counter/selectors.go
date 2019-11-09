package counter

/*
SelectCounter selects counter value
*/
func (state *ApplicationState) SelectCounter() int {
	return state.counter
}
