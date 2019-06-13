package vk

import (
	"testing"
	"time"
)

func defaultConfig() {

	tokenizer := DefaultTokenizer{}
	tokenizer.tokens = append(tokenizer.tokens, Token{Token: "0bff7559aa95dcac192ee624f63f6199bb8492d49c753bd83200d4eb76ddcf879502c4859754c880fdbf9"})

	vk = &VK{
		Token:         &tokenizer,
		DelayPerToken: 350 * time.Millisecond,
	}
	vk.Group.vk = vk
	vk.Wall.vk = vk
	vk.Like.vk = vk
}

func TestDefaultTokenizer_GetToken(t *testing.T) {

	tokenizer := DefaultTokenizer{}
	tokenizer.tokens = append(tokenizer.tokens, Token{Token: "1"})

	tkn := tokenizer.GetToken()
	if tkn.Token != "1" {
		t.Fatal("Toket != 1")
	}
}

func TestDefaultTokenizer_PutToken(t *testing.T) {

	tokenizer := DefaultTokenizer{}
	tokenizer.tokens = append(tokenizer.tokens, Token{Token: "1"})

	tkn1 := tokenizer.GetToken()
	tokenizer.PutToken(tkn1)
	tkn2 := tokenizer.GetToken()
	if tkn2.Count != 1 {
		t.Fatal("Toket.Count != 1")
	}
}
