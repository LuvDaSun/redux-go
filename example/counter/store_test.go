package counter

import (
	"sync"
	"testing"

	"github.com/LuvDaSun/redux-go/redux"
	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	const count1 = 100
	const count2 = 1000

	store := CreateApplicationStore()
	stateChannel := make(chan *ApplicationState)
	unsubscribe := store.Subscribe(func(state redux.State) {
		stateChannel <- state.(*ApplicationState)
	})
	defer unsubscribe()

	state1 := store.GetState().(*ApplicationState)

	store.DispatchChannel <- &IncrementAction{}
	state2 := <-stateChannel

	store.DispatchChannel <- &IncrementAction{}
	state3 := <-stateChannel

	store.DispatchChannel <- &DecrementAction{}
	state4 := <-stateChannel

	assert.Equal(test, 0, state1.SelectCounter())
	assert.Equal(test, 1, state2.SelectCounter())
	assert.Equal(test, 2, state3.SelectCounter())
	assert.Equal(test, 1, state4.SelectCounter())

	var wg sync.WaitGroup

	job := func() {
		for range [count1]int{} {
			store.DispatchChannel <- &IncrementAction{}
			<-stateChannel
		}
		wg.Done()
	}

	wg.Add(count2)
	for range [count2]int{} {
		go job()
	}

	wg.Wait()

	state5 := store.GetState().(*ApplicationState)
	assert.Equal(test, 1+count1*count2, state5.SelectCounter())
}
