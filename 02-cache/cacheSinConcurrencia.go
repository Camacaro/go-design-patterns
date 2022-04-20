package main

import (
	"fmt"
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
}

// Es de tipo Func que recibe un int y devuelve una interfaz y error
// Cuando usamos una interface{} es como en otros lenguaes un generico (Typescrip por ejemplo)
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
	result, exist := m.cache[key]

	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}

	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	// Un ambiente controlado por lo cual enviamos nil
	return Fibonacci(n), nil
}

/*
	Este concepto creo que tambien se llama memoization.
	https://es.wikipedia.org/wiki/Memoización

*/
func main() {
	cache := NewCache(GetFibonacci)
	fib := []int{42, 40, 41, 42, 38}
	for _, n := range fib {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", n, time.Since(start), value)
	}
}
