//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

func (p *Player) modifyHealth(newHealth int) {
	diff := newHealth - p.health
	fmt.Println("Difference is", diff)
	p.health = newHealth
}

func (p *Player) modifyEnergy(newEnergy int) {
	diff := newEnergy - p.energy
	fmt.Println("Difference is", diff)
	p.energy = newEnergy
}

func main() {
	player := Player{
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
		name:      "John",
	}
	player.modifyHealth(50)
	player.modifyEnergy(40)
}
