package cmd

import (
	"github.com/saeedmdd/go-hexa/internal/adapters/http"
	"github.com/spf13/cobra"
)


var serveCommand = &cobra.Command{
	Use: "serve",
	Short: "serves application at given port",
	Run: func(cmd *cobra.Command, args []string) {
		http.Init()
	},
}


func init(){
	rootCammand.AddCommand(serveCommand)
}