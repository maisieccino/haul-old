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
	Use:   "haul",
	Short: "Manage your Haul",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("haul")
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
			viper.SetConfigName(".haul")
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.haul.yaml)")
	rootCmd.AddCommand(serveCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
