package main

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/core"
	"time"
)

type Human struct {
	core.EmptyEntity
	age       int
	ageTicker *time.Ticker
}

func (h *Human) Startup() {
	fmt.Println("Hello,I am Startup.")
	h.ageTicker = core.StartCoroutine(time.Second, func() {
		h.age++
	})
}

type MyWorld struct {
	core.World
}

func (w *MyWorld) OnInit() {
	w.AddEntity(&Human{})
}

func main() {
	proxy := core.NewWorldProxy(&MyWorld{})
	proxy.Spin()
}
