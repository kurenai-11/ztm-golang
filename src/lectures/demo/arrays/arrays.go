package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleaned(rooms [2]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println("room", room.name, "is cleaned")
		} else {
			fmt.Println("room", room.name, "is not cleaned")
		}
	}
}

func main() {
	rooms := [...]Room{{"1", true}, {"2", true}}
	checkCleaned(rooms)
}
