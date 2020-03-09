package demoapp

import (
	"fmt"
	"sync"
)

// DemoApp - init type
type DemoApp struct {
	name string
}

var demoApp *DemoApp
var once sync.Once

func createDemoApp(name string) *DemoApp {
	return &DemoApp{
		name: name,
	}
}

// GetDemoAppInstance - instance
func GetDemoAppInstance() *DemoApp {
	once.Do(func() {
		fmt.Println("Do once function was called!")
		demoApp = createDemoApp("App by Karan!")
	})

	return demoApp
}
