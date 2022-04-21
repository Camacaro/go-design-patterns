package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection(index int) {
	fmt.Println("Creating single connection", index)
	time.Sleep(time.Second * 2)
	fmt.Println("Single connection created", index)
}

var db *Database
var lock sync.Mutex

func getDatabaseIntance(index int) *Database {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		fmt.Println("Creating new database instance:", index)
		db = &Database{}
		db.CreateSingleConnection(index)
	} else {
		fmt.Println("Returning existing database instance", index)
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			getDatabaseIntance(index + 1)
		}(i)
	}

	wg.Wait()
}
