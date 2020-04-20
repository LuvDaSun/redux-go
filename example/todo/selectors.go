package todo

/*
SelectTaskCount gets the number of tasks
*/
func (state *ApplicationState) SelectTaskCount() int {
	return state.task.items.Len()
}

/*
SelectTaskCompleteCount gets the number of completed tasks
*/
func (state *ApplicationState) SelectTaskCompleteCount() int {
	counter := 0

	walker := func(key []byte, itemRaw interface{}) bool {
		item := itemRaw.(*TaskItem)
		if item.complete {
			counter++
		}
		return false
	}

	state.task.items.Root().Walk(walker)
	return counter
}
