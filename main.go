package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func main() {
	var greeting string

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			if greeting != "" {
				fmt.Println(greeting)
			} else {
				fmt.Println(viper.GetString("greeting"))
			}
		},
	}

	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting message")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is config.yaml)")

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found...")
	}
}
