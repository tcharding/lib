// memo provides memoizing of functions
// this version uses a mutex to guard a cache (map variable) used by calling goroutines
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type result struct {
	value interface{}
	err   error
}

// New returns a memoization of f.
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e, ok := memo.cache[key]
	if !ok {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	//	fmt.Fprintf(os.Stderr, "repead request: %s\n", key)
	return e.res.value, e.res.err
}
