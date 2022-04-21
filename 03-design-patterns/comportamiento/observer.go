package main

import (
	"fmt"
	"time"
)

// Clases Principales (interfaces)
type Topic interface {
	registerObserver(o Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// END

/*
	Este es un Item que implementa la interfaz Topic

	La idea es que este Item pueda ser observado por un EmailClient
	Item -> No disponible
	Item -> Cuando este disponible:Notify -> Si hay
*/
type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) registerObserver(o Observer) {
	i.observers = append(i.observers, o)
}

/* END */

// ====== Esto es un Observer ======
type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("[Observer] EmailClient %s: Item %s is available\n", eC.id, value)
}

// ====== END ======

func main() {
	nvidiaItem := NewItem("Nvidia GTX 1080")

	firstObserver := &EmailClient{
		id: "12ab",
	}

	secondObserver := &EmailClient{
		id: "34cd",
	}

	nvidiaItem.registerObserver(firstObserver)
	nvidiaItem.registerObserver(secondObserver)

	time.Sleep(time.Second * 2)
	nvidiaItem.UpdateAvailable()

	time.Sleep(time.Second * 2)
	nvidiaItem.UpdateAvailable()

	time.Sleep(time.Second * 2)
	nvidiaItem.UpdateAvailable()
}
