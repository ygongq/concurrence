package mutex

import (
	"fmt"
	"log"
	"sync"
)

// 将mutex封装到结构体中
type Counter struct {
	mu     sync.Mutex
	countV int
}

func (c *Counter) incr() {
	// 当一个 goroutine 通过调用 Lock 方法获得了这个锁的拥有权后， 其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，
	// 直到锁被释放并且自己获取到了这个锁的拥有权。
	c.mu.Lock()
	defer c.mu.Unlock()
	// count++ 不是一个原子操作，它至少包含：读取变量当前值；对该值+1；将结果再次保存到变量中
	c.countV++
}

// 获取的时候可能不能得到刚增加的值，所以也要加锁
func (c *Counter) count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.countV
}

func GetCount(work, incrCount int) {
	if work < 1 {
		log.Fatal("工作协程个数必须大于0")
	}

	if incrCount < 1 {
		log.Fatal("自增次数必须大于0")
	}

	var c Counter

	var wg sync.WaitGroup
	wg.Add(work)

	for i := 0; i < work; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrCount; j++ {
				c.incr()
			}
		}()
	}

	wg.Wait()

	fmt.Println(c.count())
}
