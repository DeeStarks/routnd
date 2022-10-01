package process

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/deestarks/routnd/config"
	"github.com/deestarks/routnd/internal/utils"
	"github.com/spf13/viper"
)

// Process is a process
type Process struct {
	Name         string
	PID          int
	Command      string
	Log          string
	CreationTime string
}

// GetProcesses gets the processes
func GetProcesses(cfg *config.Config) []Process {
	// Get the processes
	var processes []Process

	// Get the processes
	procFile := cfg.ConfigPath + "/processes.yaml"
	viper.SetConfigFile(procFile)
	if err := viper.ReadInConfig(); err != nil {
		return processes
	}
	processesMap := viper.GetStringMap("processes")

	// Convert the processes
	for name, values := range processesMap {
		valuesMap := values.(map[string]interface{})
		processes = append(processes, Process{
			Name:         name,
			PID:          valuesMap["pid"].(int),
			Command:      valuesMap["command"].(string),
			Log:          valuesMap["log_file"].(string),
			CreationTime: valuesMap["creation_time"].(string),
		})
	}

	return processes
}

// Start starts a process
func Start(cfg *config.Config, name, command string, logFile string) {
	// Get the process
	process := getProcess(cfg, name)

	// Check if the process is running
	if process.PID != 0 {
		fmt.Printf("Process \"%s\" is already running\n", name)
		os.Exit(1)
	}

	// Open log file in the args, else use the default
	var log *os.File
	if logFile == "" {
		logFile = filepath.Join(cfg.LogPath, name+".log")
	}

	log, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Split the command
	sCmd, sArgs := utils.SplitCommand(command)
	// Start the process and log the output to a file
	cmd := exec.Command(sCmd, sArgs...)
	cmd.Stdout = log
	cmd.Stderr = log
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Save the process
	// Get log absolute path
	logFile, err = filepath.Abs(logFile)
	if err != nil {
		panic(err)
	}
	saveProcess(cfg, cmd.Process.Pid, name, command, logFile)
}

// Stop stops a process
func Stop(cfg *config.Config, name string) Process {
	// Get the process
	process := getProcess(cfg, name)

	// Check if the process is running
	if process.PID == 0 {
		fmt.Printf("Process \"%s\" is not running\n", name)
		os.Exit(1)
	}

	// Stop the process
	osProcess, err := os.FindProcess(process.PID)
	if err != nil {
		panic(err)
	}
	osProcess.Kill()

	// Remove the process
	removeProcess(cfg, name)
	fmt.Printf("Process \"%s\" stopped\n", name)
	return process
}

// Status gets the status of a process
func Status(cfg *config.Config, name string) {
	// Get the process
	process := getProcess(cfg, name)

	// Check if the process is running
	if process.PID == 0 {
		fmt.Printf("Process \"%s\" is not running\n", name)
		os.Exit(1)
	}

	// Get the status
	osProcess, err := os.FindProcess(process.PID)
	if err != nil {
		panic(err)
	}

	// Check if the process is running
	if err := osProcess.Signal(syscall.Signal(0)); err != nil {
		fmt.Printf("Process \"%s\" is not running\n", name)
		os.Exit(1)
	}

	fmt.Printf("Process \"%s\" is running\n", name)
}

// Logs views the logs of a process
func Logs(cfg *config.Config, name string, tail int, follow bool) {
	// Get the process
	process := getProcess(cfg, name)

	// Check if the process is running
	if process.PID == 0 {
		fmt.Printf("Process \"%s\" is not running\n", name)
		os.Exit(1)
	}

	// View the logs
	utils.ViewFile(process.Log, tail, follow)
}

// getProcess gets a process
func getProcess(cfg *config.Config, name string) Process {
	// Get the processes
	processes := GetProcesses(cfg)

	// Get the process
	for _, p := range processes {
		if p.Name == name {
			return p
		}
	}

	return Process{}
}

// saveProcess saves a process
func saveProcess(cfg *config.Config, pid int, name, comm, logfile string) {
	// Get the processes
	processes := GetProcesses(cfg)

	// Add the process
	processes = append(processes, Process{
		Name:         name,
		PID:          pid,
		Command:      comm,
		Log:          logfile,
		CreationTime: utils.GetCurrentTimeString(),
	})

	// Save the processes
	saveProcesses(cfg, processes)
}

// removeProcess removes a process
func removeProcess(cfg *config.Config, name string) {
	// Get the processes
	processes := GetProcesses(cfg)

	// Remove the process
	for i, p := range processes {
		if p.Name == name {
			processes = append(processes[:i], processes[i+1:]...)
		}
	}

	// Save the processes
	saveProcesses(cfg, processes)
}

// saveProcesses saves the processes
func saveProcesses(cfg *config.Config, processes []Process) {
	// Get the processes
	processesMap := map[string]map[string]interface{}{}
	for _, p := range processes {
		processesMap[p.Name] = map[string]interface{}{
			"pid":           p.PID,
			"command":       p.Command,
			"log_file":      p.Log,
			"creation_time": p.CreationTime,
		}
	}

	// Save the processes
	procFile := cfg.ConfigPath + "/processes.yaml"
	viper.Set("processes", processesMap)
	if err := viper.WriteConfigAs(procFile); err != nil {
		panic(err)
	}
}
