package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, _ := user.Current()
	fmt.Printf("Running as: %s\n", u.Name)
}
