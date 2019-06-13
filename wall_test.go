package vk

import (
	"log"
	"testing"
)

func TestWall_GetPosts(t *testing.T) {

	defaultConfig()

	posts, err := vk.Wall.GetPosts(-1)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) == 0 {
		t.Fatal("Result = 0")
	}
	//log.Printf("Posts: %#v", posts)
}

func TestWall_GetComments(t *testing.T) {

	OwnerID := -1
	ItemID := 392167

	c, err := vk.Wall.GetComments(OwnerID, ItemID)
	if err != nil {
		t.Fatal(err)
	}
	if len(c) == 0 {
		t.Fatal("Result = 0")
	}
	log.Printf("[I] Comments https://vk.com/wall%d_%d => %d\n", OwnerID, ItemID, len(c))
	//for _, cc := range c {
	//	log.Println("", cc.FromID, cc.Thread.Count, cc.Text)
	//}
}
