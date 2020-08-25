package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

const finalWord = "Go!"
const startValue = 3

const write = "write"
const sleep = "sleep"

// Sleeper is the interface for extracting the sleep functionality
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is the normal Sleeper
type DefaultSleeper struct {}

// ConfigurableSleeper can be configured regarding the time to wait
type ConfigurableSleeper struct {
    duration 	time.Duration
    sleep    	func(time.Duration)
}

// CountdownOperationsSpy spies on the Sleep Calls
type CountdownOperationsSpy struct {
    Calls []string
}

// Sleep is the mocked implementation for Sleeping
func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

// Write is the mocked implementation for writing
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

// Sleep is a mock implementation for Sleep
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// SpyTime mocks the sleeping
type SpyTime struct {
    durationSlept time.Duration
}

// Sleep is the mocked implementation
func (s *SpyTime) Sleep(duration time.Duration) {
    s.durationSlept = duration
}

// Sleep is the mocked implementation
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts down from a given number
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := startValue; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}