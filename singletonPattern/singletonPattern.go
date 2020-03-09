package main

import (
	"github.com/krnblni/GoingGo/singletonPattern/demoApp"
)

func main() {
	for i := 0; i < 10; i++ {
		demoapp.GetDemoAppInstance()
	}
}