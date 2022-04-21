package main

import "fmt"

/*
	Queremos crear una clase que implemente una interfaz de pago
*/
type Payment interface {
	Pay()
}

/*
	Tenemos la clase Cash que implementa la interfaz Payment
*/
type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying with cash")
}

/*  CashPayment END */

/*
	Tenemos la clase BankPayment que implementa la interfaz Payment

	Pero aqui tenemos una variante el Pay necesita un numero de
	cuenta entonces se le pasa por parameto el numero de cuenta.
	Ahora bien esto no cumpla con la interfaz Payment porque no
	se recibe un parametro.
*/
type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying with bank account %d\n", bankAccount)
}

/*  BankPayment END */

/*
	Se crea un adaptador para clase BankPayment
	se ponen dos propiedades BankPayment(La clase anterior) y bankAccount
*/
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

/*
	Ahora creamos una funcion llamada Pay() esta es la que
	cumplirá con la interfaz Payment y dentro de ella ejecutaremos
	la funcion Pay() de la clase BankPayment que si recibe un parametro
*/
func (b *BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.bankAccount)
}

/*  BankPaymentAdapter END */

/*
	Funcion comun para todos los metodos de pago
*/
func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	cash := CashPayment{}
	ProcessPayment(cash)

	/* BankPayment no cumple con la interfaz Payment */
	// bank := BankPayment{}
	// ProcessPayment(bank)

	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 123,
	}
	ProcessPayment(bank)
}
