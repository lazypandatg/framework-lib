package Queue

import "sync"

// Group Pool Goroutine Pool
type Group struct {
	queue chan int
	wg    *sync.WaitGroup
}

// NewGroup New 新建一个协程池
func NewGroup(size int) *Group {
	if size <= 0 {
		size = 1
	}
	return &Group{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 新增一个执行
func (p *Group) Add(delta int) {
	// delta为正数就添加
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	// delta为负数就减少
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Group) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait 等待Goroutine执行完毕
func (p *Group) Wait() {
	p.wg.Wait()
}
