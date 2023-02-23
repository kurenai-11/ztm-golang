//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package rcv

import "testing"

var testPlayer *Player

func init() {
	testPlayer = NewPlayer("Test Player")
}

func TestModifyHealth(t *testing.T) {
	testPlayer := *testPlayer
	testPlayer.modifyHealth(50)
	if testPlayer.health != 50 {
		t.Errorf("TestPlayer.health should return 50, got %v", testPlayer.health)
	}
	testPlayer.modifyHealth(200)
	if testPlayer.health != 100 {
		t.Errorf("TestPlayer.health should return 100, got %v", testPlayer.health)
	}
	testPlayer.modifyHealth(-50)
	if testPlayer.health != 0 {
		t.Errorf("TestPlayer.health should return 0, got %v", testPlayer.health)
	}
}

func TestModifyEnergy(t *testing.T) {
	testPlayer := *testPlayer
	testPlayer.modifyEnergy(50)
	if testPlayer.energy != 50 {
		t.Errorf("TestPlayer.energy should return 50, got %v", testPlayer.energy)
	}
	testPlayer.modifyEnergy(200)
	if testPlayer.energy != 100 {
		t.Errorf("TestPlayer.energy should return 100, got %v", testPlayer.energy)
	}
	testPlayer.modifyEnergy(-50)
	if testPlayer.energy != 0 {
		t.Errorf("TestPlayer.energy should return 0, got %v", testPlayer.energy)
	}
}
