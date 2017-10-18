package main

import "fmt"

//interface
type vehicle interface {
	start()
	stop()
}

//implementation
type car int

//sobald ich alle methoden vom interface vehicle implementiert habe,
//wird es zum vehicle !!! crazy!!
func (c car) start() {
	fmt.Println("Car Started")
}

func (c car) stop() {
	fmt.Println("Car Stopped")
}

//motor implementiert nur start function
type motor float64
func (m motor) start() {
	fmt.Println("Motor Started")
}

//beakage implementiert nur stop function
type breakage rune
func (b breakage) stop() {
	fmt.Println("Breakage Stopped")
}

//jetzt baue ich motor und breakage zusammen
//dann werden sie auch zum vehicle!!! crazy!!
type monster struct {
	motor
	breakage
}


func move(v vehicle) {
	v.start()
	v.stop()
}

func main() {
	var c car
	move(c)

	var m monster
	move(m)
}