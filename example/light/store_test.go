package light

import (
	"sync"
	"testing"

	"github.com/LuvDaSun/redux-go/redux"
	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	const count1 = 8
	const count2 = 16
	const count3 = 32

	store := CreateApplicationStore()
	stateChannel := make(chan *ApplicationState)
	unsubscribe := store.Subscribe(func(state redux.State) {
		stateChannel <- state.(*ApplicationState)
	})
	defer unsubscribe()

	state1 := store.GetState().(*ApplicationState)

	store.DispatchChannel <- &SwitchOnAction{}
	state2 := <-stateChannel

	store.DispatchChannel <- &SwitchOffAction{}
	state3 := <-stateChannel

	store.DispatchChannel <- &ToggleAction{}
	<-stateChannel
	state4 := <-stateChannel

	assert.Equal(test, false, state1.SelectLightIsOn())
	assert.Equal(test, true, state2.SelectLightIsOn())
	assert.Equal(test, false, state3.SelectLightIsOn())
	assert.Equal(test, true, state4.SelectLightIsOn())

	for range [count1]int{} {
		var wg sync.WaitGroup

		job := func() {
			for range [count2]int{} {
				store.DispatchChannel <- &ToggleAction{}
				<-stateChannel
				<-stateChannel
			}
			wg.Done()
		}

		wg.Add(count3)
		for range [count3]int{} {
			go job()
		}

		wg.Wait()

		state5 := store.GetState().(*ApplicationState)
		assert.Equal(test, true, state5.SelectLightIsOn())
	}
}
