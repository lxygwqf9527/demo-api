package main

import (
	"fmt"

	"github.com/lxygwqf9527/demo-api/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
