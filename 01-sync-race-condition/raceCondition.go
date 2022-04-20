package main

import "fmt"

type Balance struct {
	amount int
}

func (b *Balance) Deposit(value int) {
	b.amount = value + b.amount
}

func (b *Balance) withdraw(value int) {
	if value > b.amount {
		fmt.Println("insufficient funds")
		return
	}
	b.amount = value + b.amount
}

/*
	Race Condition:

	Aqui tenemos un problema con las goRoutina
	Estas se estaran ejecutando de manera concurrente
	y no sabremos que estara pasando con el balance puede
	ocurrir que se ejecute primero el withdraw y luego el deposit
	Ocurriendo un error de que el balance queda en -

	Basados en el código, nos devolveria un error o no?
	No podemos estar seguros, pero claramente hay un
	riesgo de que se ejecute primero el withdraw y luego el deposit
	(riesgo condicion de carrera)

	En el main.go tenemos la solucion para este problema
*/
func main() {
	balance := Balance{amount: 500}

	go balance.Deposit(100)
	go balance.Deposit(200)

	go balance.withdraw(700)
}
