package subscription

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(test *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const interval = time.Millisecond * 1000
	const count = 100

	store := CreateApplicationStore(ctx, interval)

	for index := 0; index < count; index++ {
		store.Dispatch(&StartAction{ID: index})
	}

	<-time.NewTimer(interval / 2).C

	state1 := store.GetState().(*ApplicationState)

	<-time.NewTimer(interval).C

	state2 := store.GetState().(*ApplicationState)

	<-time.NewTimer(interval).C

	state3 := store.GetState().(*ApplicationState)

	for index := 0; index < count; index++ {
		store.Dispatch(&StopAction{ID: index})
	}

	<-time.NewTimer(interval).C

	state4 := store.GetState().(*ApplicationState)

	assert.Equal(test, 0*count, state1.SelectCounter())
	assert.Equal(test, 1*count, state2.SelectCounter())
	assert.Equal(test, 2*count, state3.SelectCounter())
	assert.Equal(test, 2*count, state4.SelectCounter())
}
