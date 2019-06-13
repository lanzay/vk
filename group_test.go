package vk

import (
	"testing"
)

func TestVkGroup_GetIdByUrl(t *testing.T) {
	defaultConfig()

	//https://vk.com/umnie_knizki?w=wall-44155346_79580
	//https://vk.com/club27669892?w=wall-27669892_6721
	//https://vk.com/public44155346?w=wall-44155346_66766

	id := vk.Group.GetIdByUrl("https://vk.com/umnie_knizki?w=wall-44155346_79580")
	if id != "umnie_knizki" {
		t.Fatal("Result != exp")
	}

	id = vk.Group.GetIdByUrl("https://vk.com/club27669892?w=wall-27669892_6721")
	if id != "27669892" {
		t.Fatal("Result != exp")
	}

	id = vk.Group.GetIdByUrl("https://vk.com/public44155346?w=wall-44155346_66766")
	if id != "44155346" {
		t.Fatal("Result != exp")
	}

}
func TestGroup_GetInfo(t *testing.T) {
	defaultConfig()

	i, err := vk.Group.GetInfo("apiclub")
	if err != nil {
		t.Fatal(err)
	}
	if len(i) == 0 {
		t.Fatal("Result = 0")
	}
	if i[0].ID != 1 {
		t.Fatal("Result ID != 1")
	}
	t.Log(i)

}

func TestVkGroup_GetMembers(t *testing.T) {
	defaultConfig()

	m, err := vk.Group.GetMembers(9)
	if err != nil {
		t.Fatal(err)
	}
	if len(m) == 0 {
		t.Fatal(err)
	}
	t.Log(m)
}
