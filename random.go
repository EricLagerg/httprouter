package httprouter

import (
	"math/rand"
	"time"

	"github.com/SermoDigital/helpers"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	n           = 32
	letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// reqID returns a random hash + the timestamp in the format
// hash + ":" + time.UnixNano
func reqID() string {
	var b [1 + n + 64]byte // +1 colon

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; remain-- {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
	}
	b[n] = ':'
	return string(b[:n+1+copy(b[n+1:], helpers.FormatUint(uint64(time.Now().UnixNano())))])
}

// A pool is about 60ns per ID, while reqID itself is ~500ns.
// BenchmarkReqID-4    	 3000000	       499 ns/op
// BenchmarkReqIDPool-4	20000000	        62.0 ns/op

// randPool is a pool of random names used for rotating log files.
type randPool struct {
	c chan string
}

var pool = newRandPool(5000)

// newRandPool creates a new pool of random names and immediately
// initializes the pool with N new names.
func newRandPool(n int) *randPool {
	pool := &randPool{make(chan string, n)}
	for i := 0; i < n; i++ {
		pool.put(reqID())
	}
	return pool
}

// get gets a name from the pool, or generates a new name if none
// exist.
func (p *randPool) get() (s string) {
	select {
	case s = <-p.c:
		// get a name from the pool
	default:
		return reqID()
	}
	return
}

// put puts a new name (back) into the pool, or discards it if the pool
// is full.
func (p *randPool) put(s string) {
	select {
	case p.c <- s:
		// place back into pool
	default:
		// discard if pool is full
	}
}
