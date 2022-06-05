package main

import (
	"math/rand"
	"testing"
	"time"
)

// Ensure that concurrent Read-Write for JobId is unique
func TestGetJobId(t *testing.T) {
	const numTest = 100
	ch := make(chan uint64, numTest)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= numTest; i++ {
		go func() {
			n := rand.Intn(10)
			time.Sleep(time.Duration(n) * time.Microsecond)
			id := GetItemJobId()
			ch <- id
		}()
	}
	get := 1
	var countSort [numTest + 1]uint64
	for val := range ch {
		// t.Log(val)
		countSort[val]++
		get++
		if get > numTest {
			close(ch)
		}
	}
	// t.Log(countSort)
	for _, res := range countSort {
		// t.Log(res)
		if res > 1 {
			t.Errorf("Duplicate found")
		}
	}
}
