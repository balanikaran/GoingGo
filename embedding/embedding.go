package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this is the outer type
type myCustomRandom struct {
	*rand.Rand // embedding Rand type in my custom type
	// Rand is the inner type
	count int
}

func newMyCustomRandom(i int64) *myCustomRandom {
	return &myCustomRandom{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *myCustomRandom) getNewRandomInRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *myCustomRandom) getCount() int {
	return cr.count
}

// We can also have a function in outer type same as
// the inner type (this new function will be given priority)
// If this function is un-commented, Statement A and B will differ!

// func (cr *myCustomRandom) Intn(n int) int {
// 	fmt.Println("function of outer type was called!")
// 	cr.count ++
// 	return cr.Rand.Intn(n)
// }

func main() {
	cr := newMyCustomRandom(time.Now().UnixNano())
	fmt.Println(cr.getNewRandomInRange(3, 10))
	fmt.Println(cr.getCount())

	// We can direclty call functions of inner type
	fmt.Println(cr.Intn(10))      // Statement A
	fmt.Println(cr.Rand.Intn(10)) // Statement B
	// Statement A and B are both same (not the output though!)
}
