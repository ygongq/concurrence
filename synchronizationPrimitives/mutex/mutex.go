package mutex

import (
	"fmt"
	"sync"
)

func NotMutex() {
	var wg sync.WaitGroup
	wg.Add(10)

	var count int
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}

func Mutex() {
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)

	var count int
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				m.Lock()
				count++
				m.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
