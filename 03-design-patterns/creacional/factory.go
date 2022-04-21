package main

import "fmt"

// Clase principal (interface)
type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// ===== Computer Class =========================
/*
	Computer satisfies the IProduct interface
*/
type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

// ===== END Computer Class =========================

// ===== Laptop de tipo Product =========================
type Laptop struct {
	Computer // Composition (inheritance) - Composicion sobre la herencia
}

func newLaptop(name string) IProduct {
	return &Laptop{
		Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// ===== END Laptop de tipo Product =========================

// ===== Destock de tipo Product =========================
type Destock struct {
	Computer // Composition (inheritance) - Composicion sobre la herencia
}

// func newDestock(name string) IProduct {
// 	return &Destock{
// 		Computer: Computer{
// 			name: "Destock Computer",
// 			stock: 35,
// 		},
// 	}
// }

func newDestock(name string) IProduct {
	return &Destock{
		Computer{
			name:  "Destock Computer",
			stock: 35,
		},
	}
}

// ===== END Destock de tipo Product =========================

// ===== Factory =========================
/*
	Se encarga de la intencia de los objetos
*/
func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop("Laptop"), nil
	}

	if computerType == "destock" {
		return newDestock("Destock"), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

// ===== END Factory =========================

func printNameAndStock(p IProduct) {
	fmt.Printf("Product Name: %s, Stock: %d\n", p.getName(), p.getStock())
}

func main() {
	// Instanciamos un objeto de tipo Laptop
	laptop, _ := GetComputerFactory("laptop")
	laptop.setName("Laptop")
	laptop.setStock(10)
	printNameAndStock(laptop)

	// Instanciamos un objeto de tipo Destock
	destock, _ := GetComputerFactory("destock")
	destock.setName("Destock")
	destock.setStock(20)
	printNameAndStock(destock)
}
