package commands

import (
	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/process"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a process",
	Long:  `Start a process in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		start(cmd.Flags())
	},
}

// Add the flags
func init() {
	StartCmd.Flags().StringP("name", "n", "", "The name of the process")
	StartCmd.Flags().StringP("command", "c", "", "The command to run")
	StartCmd.Flags().StringP("log-file", "l", "", "The log file to save the output to (optional)")
}

func start(flags *pflag.FlagSet) {
	// Get the name
	name, err := flags.GetString("name")
	if err != nil {
		panic(err)
	}

	// Get the command
	command, err := flags.GetString("command")
	if err != nil {
		panic(err)
	}

	// Get log file
	logFile, err := flags.GetString("log-file")
	if err != nil {
		panic(err)
	}

	// Get the config
	cfg := config.New()

	// Start the process
	process.Start(cfg, name, command, logFile)
}
