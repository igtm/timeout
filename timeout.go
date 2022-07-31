package timeout

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

func Timeout(fn func() error, duration time.Duration) error {
	ch := make(chan error, 1)
	go func() {
		ch <- fn()
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(duration):
		return ErrTimeout
	}
}
