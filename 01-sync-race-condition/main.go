package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

// mux *sync.Mutex
func Deposit(value int, wg *sync.WaitGroup, mux *sync.RWMutex) {

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

		RWMutex:
		Se sigue usando igual ya que si necesitamos bloquear el programa,
		ya que esta escribiendo en la variable balance, la esta mutando
		pero en Balance() necesitamos leer la variable balance,
	*/
	mux.Lock()

	b := balance
	balance = value + b

	// Desbloqueamos el programa
	mux.Unlock()
}

func Balance(mux *sync.RWMutex) int {
	/*
		Aqui usamos el mux para bloquear el acceso a la variable
		balance, para que no se modifique mientras se lee
		De esta manera, no se puede modificar la variable balance

		Otras goRoutine no podran escribir en la variable balance,
		pero si leerla
	*/
	mux.RLock()
	b := balance
	mux.RLock()
	return b
}

/*
	Race Condition Solution:

	sync.Mutex.Lock() nos ayudará a bloquear el acceso a valores
	compartidos en diferentes GoRoutines

	sync.Mutex.Unlock() desbloqueará nuevamente el valor al que
	necesitamos acceder

	1 Depositando -> Escribiendo en la variable balance (Race Condition)
	N Balance() -> Leer

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
	// var mux sync.Mutex

	/*
		Aqui tenemos una variable para condicionar el acceso
		de lectura y escritura a la variable balance

		Más abajo hay una mejor descrio de RWMutex
	*/
	var mux sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &mux)
	}

	wg.Wait()
	fmt.Println("Balance: ", Balance(&mux))
}

/*
 Diferencias entre RWMutex y no usar nada:

	.Lock(): bloquea lecturas (con RLock) y escrituras (con Lock) de otras goroutines

	.Unlock(): ermite nuevas lecturas (con Rlock) y/o otra escritura (con Lock)

	.RLock(): bloquea escrituras (Lock) pero no bloquea lecturas (RLock) de otras goroutines

	.RUnlock(): permite nuevas escrituras (y también lecturas, pero por la naturaleza de RLock, estas no se vieron bloqueadas nunca)

	En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por nuevos escritores, a diferencia de no usar nada.
*/
