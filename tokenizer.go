package vk

import "time"

type Token struct {
	Token   string
	Count   int
	LastReq time.Time
	Agent   string
}

func (t *Token) GetLastReq() time.Time {
	return t.LastReq
}
func (t *Token) SetLastReq(tm time.Time) {
	t.LastReq = tm
}

//===========================================
type Tokenizer interface {
	GetToken() *Token
	PutToken(*Token)
}

type DefaultTokenizer struct {
	tokens []Token
}

func (t *DefaultTokenizer) GetToken() *Token {

	//TODO pool
	tkn := &t.tokens[0]
	tkn.LastReq = time.Now()
	return tkn
}

func (t *DefaultTokenizer) PutToken(tkn *Token) {
	//TODO pool
	tkn.Count++
}
