/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-mysqldump/internal/cli/mysqldump"
	"go-mysqldump/pkg/config"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-mysqldump",
	Short: "Mysqump Tool by Aristides",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.LoadConfig(cmd, cfgFile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log.level", "", "debug", "Log level")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./mysqldump.yaml)")

	rootCmd.AddCommand(mysqldump.NewCommand())
}
