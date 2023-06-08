package memo4

import (
	"fmt"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(string, <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key, cancel)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}

	// cancel detection
	select {
	case <-cancel:
		memo.mu.Lock()
		delete(memo.cache, key)
		memo.mu.Unlock() // delete canceled entry
		return nil, fmt.Errorf("%s is canceled", key)
	default:
	}
	return e.res.value, e.res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}
