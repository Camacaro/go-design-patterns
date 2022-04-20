package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(value int, wg *sync.WaitGroup, mux *sync.Mutex) {

	defer wg.Done()
	/*
		Lo pudiera poner aqui tambien, pero:

		Las buenas practicas dicen que solo se debe usar un defer por program
		La secuencia con mas de un defer es de la ultima a la primera
		defer wg.Done() es el ultimo defer (Se ejecuta de segundo, antes de morir el progama)
		defer lock.Unlock() es el primero (Se jecuta de primero)
	*/
	// defer lock.Unlock()

	/*
		Bloqueamos el programa hasta que se termine de ejecutar
		Esto evita que se ejecuten de manera concurrente.

		La siguiente goRoutine no se podrá ejeccutar hasta que
		se termine de ejecutar la goRoutine actual (se desbloquee)

		Esto no bloquea todas las goRoutine, solo la que esta
		ejecutando el Deposit, la que este usando la variable balance
		que es compartida por todas las goRoutine
	*/
	mux.Lock()

	b := balance
	balance = value + b

	// Desbloqueamos el programa
	mux.Unlock()
}

func Balance() int {
	b := balance
	return b
}

// func withdraw(value int) {
// 	if value > balance {
// 		fmt.Println("insufficient funds")
// 		return
// 	}
// 	balance = value + balance
// }

/*
	Race Condition Solution:

	sync.Mutex.Lock() nos ayudará a bloquear el acceso a valores
	compartidos en diferentes GoRoutines

	sync.Mutex.Unlock() desbloqueará nuevamente el valor al que
	necesitamos acceder

	Tips:
	Si quieres saber si tu codigo tiene algun data race
	debes compilarlo con el flag --race
	$ go build --race main.go
	$ /.main
	Esto te dara un warning log del race (de haber)
*/
func main() {

	var wg sync.WaitGroup
	/*
		Esto es para evitar que Deposit este usando la
		misma variable en diferentes goRoutine

		mux antes tenia el nombre de lock,
		pero es mejor renombrar lock a mux para evitar
		lo que se conoce en Go como tartamudeo lock.Lock()
	*/
	var mux sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &mux)
	}

	wg.Wait()
	fmt.Println("Balance: ", Balance())
}
