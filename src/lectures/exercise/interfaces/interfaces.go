//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	Lift()
}

type Vehicle struct {
	vehicleType string
	modelName   string
}

func (v Vehicle) Lift() {
	fmt.Println("Lifted:", v.modelName)
	switch v.vehicleType {
	case "Truck":
		fmt.Printf("Lifted the truck %v with the large lift\n", v.modelName)
	case "Car":
		fmt.Printf("Lifted the car %v with the standard lift\n", v.modelName)
	case "Motorcycle":
		fmt.Printf("Lifted the motorcycle %v with the small lift\n", v.modelName)
	}
}

func main() {
	vehicles := []Vehicle{
		{"Truck", "Grand 02 XL"},
		{"Car", "Phewer XS"},
		{"Motorcycle", "Raidle XR"},
	}
	for _, v := range vehicles {
		v.Lift()
	}
}
