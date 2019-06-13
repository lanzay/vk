package vk

import "testing"

func TestVkLike_GetList(t *testing.T) {

	defaultConfig()

	OwnerID := -1
	ItemID := 49296

	likes, err := vk.Like.GetList(OwnerID, ItemID)
	if err != nil {
		t.Fatal(err)
	}
	if len(likes) == 0 {
		t.Fatal("Result = 0")
	}
	t.Log(len(likes))
}
