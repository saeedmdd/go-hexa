package main

import (
	"github.com/saeedmdd/go-hexa/cmd"
	"github.com/saeedmdd/go-hexa/pkg/config"
)

func main() {
	config.Apply()
	cmd.Execute()
}
