package main

import (
	"fmt"
	"sync"
	"time"
)

type m  map[string]string

var mu = sync.RWMutex{}

func (m *m)set()  {
	mu.Lock()
	defer mu.Unlock()

}

func (m *m)get()  {
	mu.RLock()
	defer mu.RUnlock()
}

func main() {

	m := m{}
	for i := 0; i < 2; i++ {
		go func() {
			m["a"] = "aa"
		}()
	}
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println(m["a"])
		}()
	}
	time.Sleep(5 * time.Second)
}
