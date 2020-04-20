package todo

/*
AddTaskItemAction adds a task
*/
type AddTaskItemAction struct {
	id   string
	name string
}

/*
CompleteTaskItemAction completes a task
*/
type CompleteTaskItemAction struct {
	id string
}

/*
RemoveTaskItemAction removes a task
*/
type RemoveTaskItemAction struct {
	id string
}
