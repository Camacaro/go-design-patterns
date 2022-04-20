package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exist := m.cache[key]
	m.lock.Unlock()

	if !exist {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}

	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

/*
	Problema de concurrencia, race condition
	go run buiild --race cacheConcurrente.go
*/
func main() {
	cache := NewCache(GetFibonacci)
	fib := []int{42, 40, 41, 42, 38}

	var wg sync.WaitGroup

	for _, n := range fib {

		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", index, time.Since(start), value)

		}(n)
	}

	wg.Wait()
}
