// The `memo` package provides a secure point of view
// parallelism memorizing a function of the Func type

package memo

// Memo memorize the results of calls Func
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is type of memory function
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New creates new Memo
func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

// Get gets function cache with key
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}
