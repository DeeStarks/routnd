package commands

import (
	"os"
	"strconv"

	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/process"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// List
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the processes",
	Long:  `List all the processes that are currently running in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	// Get the config
	cfg := config.New()

	// Get the processes
	processes := process.GetProcesses(cfg)

	var data [][]string
	
	for _, p := range processes {
		data = append(data, []string{strconv.Itoa(p.PID), p.Name, p.Command, p.CreationTime})
	}

	// Create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PID", "Name", "Command", "Created"})
	// table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

// Status
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of a process",
	Long:  `Get the status of a process that is currently running in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		status(cmd.Flags())
	},
}

// Flags
func init() {
	StatusCmd.Flags().StringP("name", "n", "", "The name of the process")
}

func status(flags *pflag.FlagSet) {
	if name, err := flags.GetString("name"); err != nil {
		panic(err)
	} else {
		process.Status(config.New(), name)
	}
}
