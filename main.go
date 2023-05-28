package main

import (
	"github.com/saeedmdd/go-hexa/cmd"
	"github.com/saeedmdd/go-hexa/pkg/config"
	"fmt"
)

func main() {
	fmt.Println("akbar")
	config.Apply()
	cmd.Execute()
}
