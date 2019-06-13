package vk

import "time"

type ResponseInt struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type City struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Link struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

type APDEX struct {
	Begin     time.Time
	Wait      time.Duration
	BeginReq  time.Time
	Req       time.Duration
	RoundTrip time.Duration
}

func NewAPDEX() *APDEX {
	a := APDEX{}
	a.Start()
	return &a
}
func (a *APDEX) Round() time.Duration {
	a.RoundTrip = time.Since(a.Begin)
	a.Start()
	return a.RoundTrip
}
func (a *APDEX) ReqStart() {
	a.BeginReq = time.Now()
}
func (a *APDEX) ReqEnd() time.Duration {
	a.Req = time.Since(a.BeginReq)
	return a.Req
}
