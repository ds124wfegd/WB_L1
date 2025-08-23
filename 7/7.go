// go run -race 7.go - использовать для проверки на dataRace
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap - потокобезопасная map
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap создает новую потокобезопасную map
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set безопасно устанавливает значение
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get безопасно получает значение
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

// Delete безопасно удаляет значение
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Len безопасно возвращает количество элементов
func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// GetAll безопасно возвращает все данные
func (sm *SafeMap) GetAll() map[string]int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Возвращаем копию для безопасности
	result := make(map[string]int)
	for k, v := range sm.data {
		result[k] = v
	}
	return result
}

func main() {

	// ---------------создание потоко-безоппасной мапы---------------
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// Запись из нескольких горутин
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("id_%d", id)
			safeMap.Set(key, id*10)
		}(i)
	}

	// Чтение из нескольких горутин
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("id_%d", id%100)
			if value, exists := safeMap.Get(key); exists {
				_ = value // Используем значение
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Всего элементов в созданной потокобезопасной мапе %d\n", safeMap.Len())

	// ---------------использование встроенной sync.map---------------

	var safeMapembedded sync.Map

	for i := 0; i < 15; i++ { // Запись данных
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("id-%d-%d", id, j)
				safeMapembedded.Store(key, j*id)
				time.Sleep(time.Microsecond * 5)
			}
		}(i)
	}

	for i := 0; i < 10; i++ { // Чтение и обновление
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 30; j++ {
				key := fmt.Sprintf("id-%d-%d", id%15, j)

				safeMapembedded.LoadOrStore(key, 0) // Атомарное обновление
				safeMapembedded.Load(key)

				time.Sleep(time.Microsecond * 10)
			}
		}(i)
	}

	for i := 0; i < 5; i++ { // Удаление
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				key := fmt.Sprintf("id-%d-%d", id, j)
				safeMap.Delete(key)
				time.Sleep(time.Microsecond * 12)
			}
		}(i)
	}

	wg.Wait()

	count := 0
	// Подсчет элементов
	safeMapembedded.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	fmt.Printf("Всегоэлементов в sync.Map: %d\n", count)
}
