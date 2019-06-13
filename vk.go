package vk

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
	"time"
)

var vk *VK

type VK struct {
	DelayPerToken time.Duration
	Token         Tokenizer
	Group         VkGroup
	Wall          VkWall
	Like          VkLike
}

func (v *VK) GetBody(p map[string]string, apdex *APDEX) ([]byte, error) {

	if apdex == nil {
		apdex = NewAPDEX()
	}
	u, token := vk.UrlCompile(p)

	var err error
	var body []byte

	apdex.Wait = (vk.DelayPerToken) - time.Now().Sub(token.GetLastReq())
	<-time.After(apdex.Wait)

	token.SetLastReq(time.Now())
	apdex.ReqStart()
	_, body, err = fasthttp.Get(body, u) //<<<<<-------------
	apdex.ReqEnd()
	log.Println("[D] APDEX", apdex.Wait, apdex.Req)
	vk.Token.PutToken(token)
	if err != nil {
		log.Panicln("[F]", err)
	}

	if bytes.EqualFold([]byte("error"), body[2:7]) {
		body = body[9 : len(body)-1]
		log.Println("[I] ERR", string(body))
		return body, errors.New("VK Err")
	} else {
		//log.Println("[I] Ok")
		body = body[12 : len(body)-1]
	}
	return body, nil
}

func (v *VK) UrlCompile(p map[string]string) (string, *Token) {

	var par string
	for k, v := range p {
		if strings.EqualFold("method", k) || strings.EqualFold("v", k) {
			continue
		}
		par += fmt.Sprintf("&%s=%s", k, v)
	}

	token := vk.Token.GetToken()

	u := fmt.Sprintf("%s/%s?v=%s%s&access_token=%s", ENDPOINT, p["method"], p["v"], par, token.Token)

	return u, token
}
