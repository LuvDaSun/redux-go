package todo

import (
	"fmt"
	"sync"
	"testing"

	"github.com/LuvDaSun/redux-go/redux"
	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	const count = 1000

	store := CreateApplicationStore()
	stateChannel := make(chan *ApplicationState)
	unsubscribe := store.Subscribe(func(state redux.State) {
		stateChannel <- state.(*ApplicationState)
	})
	defer unsubscribe()

	state1 := store.GetState().(*ApplicationState)

	{
		var wg sync.WaitGroup
		job := func(index int) {
			store.DispatchChannel <- &AddTaskItemAction{
				id:   fmt.Sprintf("%d", index),
				name: string(index),
			}
			<-stateChannel
			wg.Done()
		}
		wg.Add(count)
		for index := range [count]int{} {
			go job(index)
		}
		wg.Wait()
	}
	state2 := store.GetState().(*ApplicationState)

	{
		var wg sync.WaitGroup
		job := func(index int) {
			store.DispatchChannel <- &CompleteTaskItemAction{
				id: fmt.Sprintf("%d", index),
			}
			<-stateChannel
			wg.Done()
		}
		wg.Add(count)
		for index := range [count]int{} {
			go job(index)
		}
		wg.Wait()
	}
	state3 := store.GetState().(*ApplicationState)

	{
		var wg sync.WaitGroup
		job := func(index int) {
			store.DispatchChannel <- &RemoveTaskItemAction{
				id: fmt.Sprintf("%d", index),
			}
			<-stateChannel
			wg.Done()
		}
		wg.Add(count)
		for index := range [count]int{} {
			go job(index)
		}
		wg.Wait()
	}
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, 0, state1.SelectTaskCount())
	assert.Equal(test, 0, state1.SelectTaskCompleteCount())

	assert.Equal(test, count, state2.SelectTaskCount())
	assert.Equal(test, 0, state2.SelectTaskCompleteCount())

	assert.Equal(test, count, state3.SelectTaskCount())
	assert.Equal(test, count, state3.SelectTaskCompleteCount())

	assert.Equal(test, 0, state4.SelectTaskCount())
	assert.Equal(test, 0, state4.SelectTaskCompleteCount())
}
