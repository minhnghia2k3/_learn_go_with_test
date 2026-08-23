package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const startNum = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Counter(writer io.Writer, sleeper Sleeper) {
	for i := startNum; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(writer, finalWord)
}

// func main() {
// 	// sleeper := &DefaultSleeper{}
// 	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
// 	Counter(os.Stdout, sleeper)
// }
