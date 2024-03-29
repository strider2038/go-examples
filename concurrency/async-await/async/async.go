// Package async is async/await pattern instrument based on golang typed parameters (generics).
// See https://medium.com/@jon_43067/go-generics-and-concurrency-d0dccab73a73
package async

import (
	"context"
	"sync"
)

type Func[In, Out any] func(context.Context, In) (Out, error)

type Promise[V any] struct {
	value V
	err   error
	done  <-chan struct{}
}

func (p *Promise[V]) Wait() error {
	<-p.done

	return p.err
}

func (p *Promise[V]) Get() (V, error) {
	<-p.done

	return p.value, p.err
}

func Go[In, Out any](ctx context.Context, f Func[In, Out], input In) *Promise[Out] {
	done := make(chan struct{})
	p := Promise[Out]{done: done}
	go func() {
		defer close(done)
		p.value, p.err = f(ctx, input)
	}()

	return &p
}

type Waiter interface {
	Wait() error
}

func Await(waiters ...Waiter) error {
	var wg sync.WaitGroup
	wg.Add(len(waiters))
	errors := make(chan error, len(waiters))
	done := make(chan struct{})

	for _, w := range waiters {
		go func(w Waiter) {
			defer wg.Done()
			err := w.Wait()
			if err != nil {
				errors <- err
			}
		}(w)
	}

	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case err := <-errors:
		return err
	case <-done:
	}

	return nil
}

func WithCancel[In, Out any](f Func[In, Out]) Func[In, Out] {
	return func(ctx context.Context, in In) (Out, error) {
		done := make(chan struct{})
		var out Out
		var err error
		go func() {
			defer close(done)
			out, err = f(ctx, in)
		}()

		select {
		case <-ctx.Done():
			var zero Out
			return zero, ctx.Err()
		case <-done:
		}

		return out, err
	}
}

func Then[In, Out any](ctx context.Context, p *Promise[In], f Func[In, Out]) *Promise[Out] {
	done := make(chan struct{})
	out := Promise[Out]{done: done}

	go func() {
		defer close(done)
		val, err := p.Get()
		if err != nil {
			out.err = err
			return
		}
		out.value, out.err = f(ctx, val)
	}()

	return &out
}

type CancelingWaiter struct {
	cancel func()
}

func NewCancelingWaiter(ctx context.Context) (context.Context, *CancelingWaiter) {
	ctx, cancel := context.WithCancel(ctx)

	return ctx, &CancelingWaiter{cancel: cancel}
}

func (w *CancelingWaiter) Await(waiters ...Waiter) error {
	err := Await(waiters...)
	if err != nil {
		w.cancel()

		return err
	}

	return nil
}
