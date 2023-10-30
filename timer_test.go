package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Waktu sekarang", time.Now())

	time := <-timer.C
	fmt.Println("Waktu sekarang", time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Waktu sekarang", time.Now())

	time := <-channel
	fmt.Println("Waktu sekarang", time)
}

func TestAfterFunc(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Waktu sekarang", time.Now())
		group.Done()
	})
	fmt.Println("Waktu sekarang", time.Now())

	group.Wait()
}
