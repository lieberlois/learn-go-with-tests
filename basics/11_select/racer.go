package racer

import (
	"net/http"
	"time"
	"fmt"
)

const tenSecondTimeout = 10 * time.Second


// Racer returns a new ConfigurableRacer
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer starts two HTTP Requests and returns the one with the faster answer with a configured timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// http requests now run concurrently
	// struct{} is the smallest datatype
    ch := make(chan struct{})
    go func() {
        http.Get(url)
        close(ch)
    }()
    return ch
}


// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)
	
// 	if aDuration < bDuration {
// 		return a
// 	} 
// 	return b
// }

// func measureResponseTime(url string) time.Duration {
//     start := time.Now()
//     http.Get(url)
//     return time.Since(start)
// }