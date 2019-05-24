package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addClientFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("url", "u", "", "URL to connect to")
	viper.SetDefault("client.url", "http://localhost:4285")
	viper.BindPFlag("client.url", cmd.PersistentFlags().Lookup("url"))
}
