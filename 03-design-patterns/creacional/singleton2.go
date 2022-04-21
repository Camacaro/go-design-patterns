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

/*
	tambien se puede usar singleton con el paquete sync.Once
*/
var once sync.Once

func getDatabaseIntance(index int) *Database {
	/*
		y ya este paquete te garantiza que de verdad solo se
		ejecutara una sola vez, al ser un proceso atómico
	*/
	once.Do(func() {
		fmt.Println("Creating DB connection", index)
		db = &Database{}
		db.CreateSingleConnection(index)
	})

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
