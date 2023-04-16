package main

import (
	"fmt"
	"gitee.com/ihezebin/quick-create-project/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println("\nErr happened:", err)
	}
}
