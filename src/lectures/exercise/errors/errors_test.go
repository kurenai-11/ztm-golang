package timeparse

import "testing"

func TestString(t *testing.T) {
	time := Time{}.New()
	res := time.String()
	if res != "00:00:00" {
		t.Errorf("time.String() should be 00:00:00, but got %s", res)
	}
	time.Hour = 23
	time.Minute = 59
	time.Second = 59
	res = time.String()
	if res != "23:59:59" {
		t.Errorf("time.String() should be 23:59:59, but got %s", res)
	}
}

func TestParse(t *testing.T) {
	time1 := Time{Original: "14:07:33"}
	err := time1.Parse()
	if err != nil {
		t.Errorf("time1.Parse() should be nil, but got %s", err)
	}
	if time1.String() != time1.Original {
		t.Errorf("time1.String() should be 14:07:33, but got %s", time1.String())
	}
	time2 := Time{}
	err = time2.Parse()
	if err == nil {
		t.Errorf("time2.Parse() should error because it is an empty string")
	}
	time3 := Time{Original: "25:33:17"}
	err = time3.Parse()
	if err == nil {
		t.Errorf("time3.Parse() should error because it is invalid time")
	}
}
