package main

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile = ""
)

var rootCmd = &cobra.Command{
	Use:   "haul-srv",
	Short: "Start Haul server",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.Get("bind-host")
		port := viper.Get("port")
		fmt.Printf("Starting Haul on %s:%s\n", host, port)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			viper.AddConfigPath(home)
			viper.SetConfigName(".haul-srv")
		}
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				fmt.Fprintln(os.Stderr, "Error reading in config: ", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.haul-srv.yaml)")
	rootCmd.PersistentFlags().StringP("bind-host", "b", "", "IP to bind server to")
	rootCmd.PersistentFlags().StringP("port", "p", "", "Port to listen on")
	viper.SetDefault("bind-host", "0.0.0.0")
	viper.SetDefault("port", "4285")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
