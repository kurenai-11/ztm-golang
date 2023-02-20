//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

func printStatus(status int) string {
	switch status {
	case 0:
		return "Online"
	case 1:
		return "Offline"
	case 2:
		return "Maintenance"
	case 3:
		return "Retired"
	default:
		return "Unknown"
	}
}

func printGeneralStatuses(serverStatuses map[string]int) {
	for name, status := range serverStatuses {
		fmt.Printf("%s: %v\n", name, printStatus(status))
	}
}

func displayServerStatus(serverStatuses map[string]int) {
	if len(serverStatuses) == 0 {
		return
	}
	fmt.Printf("servers: %v\n", len(serverStatuses))
	printGeneralStatuses(serverStatuses)
}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatuses := make(map[string]int)
	for _, server := range servers {
		serverStatuses[server] = Online
	}
	displayServerStatus(serverStatuses)
	serverStatuses["darkstar"] = 3
	serverStatuses["aiur"] = 1
	displayServerStatus(serverStatuses)
	for name := range serverStatuses {
		serverStatuses[name] = 2
	}
	displayServerStatus(serverStatuses)
}
