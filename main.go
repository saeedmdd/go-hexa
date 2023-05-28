package main

import (
	"fmt"
	"github.com/saeedmdd/go-hexa/cmd"
	"github.com/saeedmdd/go-hexa/pkg/config"
)

func main() {
	fmt.Println("akbar", "asghar")
	config.Apply()
	cmd.Execute()
}
