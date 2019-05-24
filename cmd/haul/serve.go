package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serveCmd.PersistentFlags().StringP("bind-host", "b", "", "IP address to bind server to")
	serveCmd.PersistentFlags().StringP("port", "p", "", "Port to listen on")
	viper.SetDefault("server.bind-host", "0.0.0.0")
	viper.SetDefault("server.port", "4285")
	viper.BindPFlags(serveCmd.PersistentFlags())
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Haul server",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.Get("server.bind-host")
		port := viper.Get("server.port")
		fmt.Printf("Starting Haul on %s:%s\n", host, port)
	},
}
