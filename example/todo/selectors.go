package todo

/*
SelectTaskCount gets the number of tasks
*/
func (state *ApplicationState) SelectTaskCount() int {
	return len(state.task.taskMap)
}

/*
SelectTaskCompleteCount gets the number of completed tasks
*/
func (state *ApplicationState) SelectTaskCompleteCount() int {
	counter := 0
	for _, v := range state.task.taskMap {
		if !v.complete {
			continue
		}
		counter++
	}
	return counter
}
