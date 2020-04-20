package todo

import (
	"github.com/LuvDaSun/redux-go/redux"
	iradix "github.com/hashicorp/go-immutable-radix"
)

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
	items *iradix.Tree
}

/*
InitialTaskState initial todo state
*/
var InitialTaskState = &TaskState{
	items: iradix.New(),
}

/*
ReduceTaskState reduces todo state
*/
func ReduceTaskState(state *TaskState, action redux.Action) *TaskState {
	switch a := action.(type) {
	case *AddTaskItemAction:
		key := []byte(a.id)
		item := &TaskItem{
			name:     a.name,
			complete: false,
		}
		items, _, exists := state.items.Insert(key, item)
		if exists {
			println(a.id)
			panic("task existed")
		}
		return &TaskState{
			items: items,
		}

	case *CompleteTaskItemAction:
		key := []byte(a.id)
		itemRaw, exists := state.items.Get(key)
		if !exists {
			panic("task did not existed")
		}
		item := itemRaw.(*TaskItem)
		items, _, exists := state.items.Insert(key, &TaskItem{
			name:     item.name,
			complete: true,
		})
		if !exists {
			panic("task did not existed")
		}
		return &TaskState{
			items: items,
		}

	case *RemoveTaskItemAction:
		key := []byte(a.id)
		items, _, exists := state.items.Delete(key)
		if !exists {
			panic("task dit not existed")
		}
		return &TaskState{
			items: items,
		}
	}

	return state
}
