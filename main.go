package main

import (
	"fmt"
	"os"

	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "routnd",
	Short: "Routnd runs processes in daemon mode",
	Long:  `Routnd is a command line tool for running processes in the background.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Version
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of routnd",
		Long:  `All software has versions. This is routnd's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.New().Version)
		},
	})
	// Start
	rootCmd.AddCommand(commands.StartCmd)
	// Stop
	rootCmd.AddCommand(commands.StopCmd)
	// Logs
	rootCmd.AddCommand(commands.LogsCmd)
	// List
	rootCmd.AddCommand(commands.ListCmd)
	// Status
	rootCmd.AddCommand(commands.StatusCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
