package main

import (
	"fmt"
	"sync"
)

// Counter с защитой через Mutex
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Запускаем 1000 горутин для инкрементации
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}

/*
//через "sync/atomic"

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// AtomicCounter - счетчик с атомарными операциями
type AtomicCounter struct {
	value int64
}

// Increment - атомарно увеличивает счетчик на 1
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Value - атомарно возвращает текущее значение счетчика
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := AtomicCounter{}

	// количество горутин для тестирования
	const numGoroutines = 1000

	// запускаем горутины для конкурентного инкремента
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// ожидаем завершения всех горутин
	wg.Wait()

	// выводим итоговое значение
	fmt.Printf("итоговое значение счетчика: %d\n", counter.Value())
	fmt.Printf("ожидаемое значение: %d\n", numGoroutines)
}
*/
