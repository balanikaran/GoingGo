// A simple Go program to fetch quote of the day
// from TheySaidSo Website

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	response, error := http.Get("http://api.theysaidso.com/qod.json")
	if error != nil {
		fmt.Println(error)
	}
	defer response.Body.Close()
	contents, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(contents))
}