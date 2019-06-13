package vk

import "time"

func (a *APDEX) Start() {
	a.Begin = time.Now()
}
