package commands

import (
	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/process"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get the logs of a process",
	Long:  `Get the logs of a process that is currently running in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		logs(cmd.Flags())
	},
}

// Flags
func init() {
	LogsCmd.Flags().StringP("name", "n", "", "The name of the process")
	LogsCmd.Flags().BoolP("follow", "f", false, "Follow the logs")
	LogsCmd.Flags().IntP("tail", "t", 0, "Number of the last lines to show")
}

func logs(flags *pflag.FlagSet) {
	var (
		name   string
		follow bool
		tail   int
		err    error
	)

	// Get the flags
	if name, err = flags.GetString("name"); err != nil {
		panic(err)
	}
	if follow, err = flags.GetBool("follow"); err != nil {
		panic(err)
	}
	if tail, err = flags.GetInt("tail"); err != nil {
		panic(err)
	}

	// Get the config
	cfg := config.New()

	// View the logs
	process.Logs(cfg, name, tail, follow)
}
