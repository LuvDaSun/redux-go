package todo

import "github.com/LuvDaSun/redux-go/redux"

/*
ApplicationState state
*/
type ApplicationState struct {
	task *TaskState
}

/*
InitialApplicationState initial application state
*/
var InitialApplicationState = &ApplicationState{
	task: InitialTaskState,
}

/*
ReduceApplicationState reduces application state
*/
func ReduceApplicationState(state *ApplicationState, action redux.Action) *ApplicationState {
	return &ApplicationState{
		task: ReduceTaskState(state.task, action),
	}
}

/*
TaskItem is the actual task
*/
type TaskItem struct {
	name     string
	complete bool
}

/*
TaskState state
*/
type TaskState struct {
	taskMap map[string]*TaskItem
}

/*
InitialTaskState initial todo state
*/
var InitialTaskState = &TaskState{
	taskMap: make(map[string]*TaskItem),
}

/*
ReduceTaskState reduces todo state
*/
func ReduceTaskState(state *TaskState, action redux.Action) *TaskState {
	switch a := action.(type) {
	case *AddTaskItemAction:
		nextState := &TaskState{
			taskMap: make(map[string]*TaskItem, len(state.taskMap)+1),
		}
		for k, v := range state.taskMap {
			nextState.taskMap[k] = v
		}
		nextState.taskMap[a.id] = &TaskItem{
			name:     a.name,
			complete: false,
		}
		return nextState

	case *CompleteTaskItemAction:
		nextState := &TaskState{
			taskMap: make(map[string]*TaskItem, len(state.taskMap)),
		}
		for k, v := range state.taskMap {
			if k == a.id {
				nextState.taskMap[k] = &TaskItem{
					name:     v.name,
					complete: true,
				}
				continue
			}
			nextState.taskMap[k] = v
		}
		return nextState

	case *RemoveTaskItemAction:
		nextState := &TaskState{
			taskMap: make(map[string]*TaskItem, len(state.taskMap)-1),
		}
		for k, v := range state.taskMap {
			if k == a.id {
				continue
			}
			nextState.taskMap[k] = v
		}
		return nextState
	}

	return state
}
