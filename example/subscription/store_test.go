package subscription

import (
	"context"
	"testing"
	"time"

	"github.com/LuvDaSun/redux-go/redux"
	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const interval = time.Millisecond * 100
	const count = 100

	store := CreateApplicationStore(ctx, interval)
	stateChannel := make(chan *ApplicationState)
	unsubscribe := store.Subscribe(func(state redux.State) {
		stateChannel <- state.(*ApplicationState)
	})
	defer unsubscribe()

	for index := 0; index < count; index++ {
		store.DispatchChannel <- &StartAction{ID: index}
		<-stateChannel
	}
	state1 := store.GetState().(*ApplicationState)

	for index := 0; index < count; index++ {
		<-stateChannel
	}
	state2 := store.GetState().(*ApplicationState)

	for index := 0; index < count; index++ {
		<-stateChannel
	}
	state3 := store.GetState().(*ApplicationState)

	for index := 0; index < count; index++ {
		store.DispatchChannel <- &StopAction{ID: index}
		<-stateChannel
	}
	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, 0*count, state1.SelectCounter())
	assert.Equal(test, 1*count, state2.SelectCounter())
	assert.Equal(test, 2*count, state3.SelectCounter())
	assert.Equal(test, 2*count, state4.SelectCounter())
}
