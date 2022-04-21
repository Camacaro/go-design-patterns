package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Reutilización de computación intensiva
*/

func ExpensiveFibonaaci(n int) int {
	fmt.Printf("Calculate Expensive Fibonnacci(%d)\n", n)
	time.Sleep(time.Second * 5)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	t := time.Now()

	s.Lock.RLock()

	exist := s.InProgress[job]

	if exist {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for response for job %d - Time: [%s]\n", job, time.Since(t))
		resp := <-response
		fmt.Printf("Got response for job %d - Time: [%s]\n", resp, time.Since(t))
		return
	}

	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()
	fmt.Printf("Working on job %d - Time: [%s]\n", job, time.Since(t))
	result := ExpensiveFibonaaci(job)

	s.Lock.RLock()
	pendingWorkers, exist := s.IsPending[job]
	s.Lock.RUnlock()

	if exist {
		for _, worker := range pendingWorkers {
			worker <- result // Le enviamos el resultado al worker
		}

		fmt.Printf("Result sent - all pending workers for job %d - Time: [%s]\n", job, time.Since(t))
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}

	wg.Wait()
}
