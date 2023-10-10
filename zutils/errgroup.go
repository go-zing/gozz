package zutils

import (
	"sync"
)

// ErrGroup is a simple version of golang.org/x/sync/errgroup.Group
type ErrGroup struct {
	wait sync.WaitGroup
	once sync.Once
	err  error
}

func (g *ErrGroup) Wait() error { g.wait.Wait(); return g.err }

func (g *ErrGroup) Go(f func() error) {
	g.wait.Add(1)
	go func() {
		defer g.wait.Done()
		if err := f(); err != nil {
			g.once.Do(func() { g.err = err })
		}
	}()
}
