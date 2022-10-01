package commands

import (
	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/process"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a process",
	Long:  `Stop a process that is currently running in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		stop(cmd.Flags())
	},
}

// Flags
func init() {
	StopCmd.Flags().StringP("name", "n", "", "The name of the process")
}

func stop(flags *pflag.FlagSet) {
	if name, err := flags.GetString("name"); err != nil {
		panic(err)
	} else {
		process.Stop(config.New(), name)
	}
}
