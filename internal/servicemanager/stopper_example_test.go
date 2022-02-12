package servicemanager

import (
	"context"
)

func ExampleStopper() {
	var s Stopper

	ready := make(chan struct{})
	go func() {
		done, ctx := s.Start(context.Background())
		defer done()

		close(ready)
		// some blocking task eg. net listener
		<-ctx.Done()
	}()
	<-ready

	// wait for task to end
	<-s.Stop()
}
