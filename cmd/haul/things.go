package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	addClientFlags(thingsCmd)
	rootCmd.AddCommand(thingsCmd)
}

var thingsCmd = &cobra.Command{
	Use:   "things",
	Short: "Manage things",
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.Get("client.url")
		fmt.Printf("Connecting to Haul on %s\n", url)
	},
}
