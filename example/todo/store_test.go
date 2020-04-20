package todo

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	const count = 10000

	store := CreateApplicationStore()
	store.Dispatch(nil)

	state1 := store.GetState().(*ApplicationState)

	{
		var wg sync.WaitGroup
		job := func(index int) {
			store.Dispatch(&AddTaskItemAction{
				id:   fmt.Sprintf("%d", index),
				name: string(index),
			})
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
			store.Dispatch(&CompleteTaskItemAction{
				id: fmt.Sprintf("%d", index),
			})
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
			store.Dispatch(&RemoveTaskItemAction{
				id: fmt.Sprintf("%d", index),
			})
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
