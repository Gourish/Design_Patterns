package main

import "fmt"

type TemperatureData struct {
	observers   map[string]Observer
	temperature float64
}

type weatherStation interface {
	Add(Observer)
	Remove(Observer)
	notifyAll()
	addData()
	getData() float64
}

func initTemperatureData() *TemperatureData {
	return &TemperatureData{observers: make(map[string]Observer), temperature: 98.4}
}

func (td *TemperatureData) Add(ob Observer) {
	td.observers[ob.getName()] = ob
}

func (td *TemperatureData) Remove(ob Observer) {
	delete(td.observers, ob.getName())
}

func (td *TemperatureData) notifyAll() {
	for _, ob := range td.observers {
		ob.update()
	}
}

func (td *TemperatureData) addData() {
	td.temperature += 0.5
	td.notifyAll()
}

func (td *TemperatureData) getData() float64 {
	return td.temperature
}

type display struct {
	ws   weatherStation
	name string
}

type Observer interface {
	update()
	getName() string
}

func initDisplay(name string, ws weatherStation) *display {
	return &display{
		name: name,
		ws:   ws,
	}
}

func (dp *display) update() {
	fmt.Printf("%s Display : current temperature %f\n", dp.name, dp.ws.getData())
}

func (dp *display) getName() string {
	return dp.name
}
