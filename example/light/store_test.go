package light

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	const count1 = 10
	const count2 = 100
	const count3 = 1000

	store := CreateApplicationStore()
	store.Dispatch(nil)

	state1 := store.GetState().(*ApplicationState)

	store.Dispatch(&SwitchOnAction{})
	state2 := store.GetState().(*ApplicationState)

	store.Dispatch(&SwitchOffAction{})
	state3 := store.GetState().(*ApplicationState)

	store.Dispatch(&ToggleAction{})
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, false, state1.SelectLightIsOn())
	assert.Equal(test, true, state2.SelectLightIsOn())
	assert.Equal(test, false, state3.SelectLightIsOn())
	assert.Equal(test, true, state4.SelectLightIsOn())

	for range [count1]int{} {
		var wg sync.WaitGroup

		job := func() {
			for range [count2]int{} {
				store.Dispatch(&ToggleAction{})
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
