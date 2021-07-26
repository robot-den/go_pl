package memo_hash

import "sync"

type Memo struct {
	f 	  Func
	cache map[string]*entry
	mu    sync.Mutex
}

type Func func(url string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{
		f: f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(url string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[url]
	if e == nil {
		ready := make(chan struct{})
		e = &entry{ready: ready}
		memo.cache[url] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(url)
		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}
	return e.res.value, e.res.err
}
