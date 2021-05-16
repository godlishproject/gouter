package gouter

import (
	"fmt"
	"net/http"
)

/**
 * let's chose the only right func to handle the incoming path
 */
func Match(path []string) {
	fmt.Println(path)
}

/**
 * just pass this func to path "/", we can handle anything here!
 */
func Handler(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("hello, " + req.URL.Query().Get("name")))
	if err != nil {
		panic(err.Error())
	}
}
