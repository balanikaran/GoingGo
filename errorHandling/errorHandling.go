package main

import (
	"fmt"
	"errors"
	"time"
)

// using a error variable
// this makes it exportable (if name starts with capital letter)
var errJoeyNotFound = errors.New("joey was not found")

// some custom error type
type errFindingJoey struct {
	msg, info string
}

// this makes errFindingJoey a child of error
func (e errFindingJoey) Error() string {
	return e.msg
}

func createSomeShit() (int, error){
	// simulate processing shit
	time.Sleep(2 * time.Second)

	// sample dictionary
	ids := map[string]int {
		"chandler": 1,
		"ross": 2,
	}

	// find joey in ids
	if _, isFound := ids["joey"]; !isFound {
		return -1, errors.New("joey was not found")
	}else {
		// never gets here!
		return 1, nil
	}

}

func createSomeShitUsingVar() (int, error) {
	// simulate processing shit
	time.Sleep(2 * time.Second)

	// sample dictionary
	ids := map[string]int {
		"chandler": 1,
		"ross": 2,
	}

	// find joey in ids
	if _, isFound := ids["joey"]; !isFound {
		return -1, errJoeyNotFound
	}else {
		// never gets here!
		return 1, nil
	}

}

func createSomeShitWithFormatting() (int, error) {
	// simulate processing shit
	time.Sleep(2 * time.Second)

	// sample dictionary
	ids := map[string]int {
		"chandler": 1,
		"ross": 2,
	}

	// find joey in ids
	if _, isFound := ids["joey"]; !isFound {
		info := "some more info about error"
		return -1, fmt.Errorf("joey was not found - %s", info)
	}else {
		// never gets here!
		return 1, nil
	}

}

func createSomeShitCustomError() (int, error) {
	// simulate processing shit
	time.Sleep(2 * time.Second)

	// sample dictionary
	ids := map[string]int {
		"chandler": 1,
		"ross": 2,
	}

	// find joey in ids
	if _, isFound := ids["joey"]; !isFound {
		return -1, errFindingJoey{"joey not found", "some info"}
	}else {
		// never gets here!
		return 1, nil
	}

}

func main() {

	if _, err := createSomeShit(); err != nil {
		// some error occured
		fmt.Println("createSomeShit() returned this: ", err)
	}else {
		// never gets here
		fmt.Println("createSomeShit() ran into success!")
	}

	if _, err := createSomeShitUsingVar(); err != nil {
		// some error occured
		fmt.Println("createSomeShitUsingVar() returned this: ", err)
	}else {
		// never gets here
		fmt.Println("createSomeShitUsingVar() ran into success!")
	}

	if _, err := createSomeShitWithFormatting(); err != nil {
		// some error occured
		fmt.Println("createSomeShitWithFormatting() returned this: ", err)
	}else {
		// never gets here
		fmt.Println("createSomeShitWithFormatting() ran into success!")
	}

	if _, err := createSomeShitCustomError(); err != nil {
		// some error occured
		fmt.Println("createSomeShitCustomError() returned this: ", err)
		if errVal, ok := err.(errFindingJoey); ok {
			fmt.Println("Other info: ", errVal.info)
		}
	}else {
		// never gets here
		fmt.Println("createSomeShitCustomError() ran into success!")
	}
	
}