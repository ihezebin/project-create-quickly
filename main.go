package main

import (
	"fmt"
	"github.com/ihezebin/project-create-quickly/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println("\nErr happened:", err)
	}
}
