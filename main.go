package main

import (
	"fmt"
	c "twenv/config"
	r "twenv/router"
)

func main() {
	// Initialize configs
	err := c.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Initialize the router
	r.Initialize()
}
