package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

// SpySleeper /* SpySleeper is used by test in order not to wait each 3 seconds, it only
type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type SpyCoundownOperations struct {
	Calls []string
}

func (s *SpyCoundownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCoundownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
