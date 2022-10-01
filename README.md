# ROUTND

Routnd (pronounced "routine-dee") is a command line tool developed in Go, for running commands/processes in the background.

## Why routnd?
I wanted to be able to run several processes with the ability to quickly stop and start them without having to open numerous terminals, and also to easily see the output of each process.

## Installation

### cURL
```bash
curl -s https://raw.githubusercontent.com/deestarks/routnd/master/scripts/install.sh | bash
```

### Manually
Download the latest release from the [releases page](https://github.com/deestarks/routnd/releases) and extract the binary to a location in your PATH.

## Usage
```bash
routnd [command] [flags]
```

### Available Commands
```bash
  help        Help about any command
  list        List all running processes
  logs        Show logs for a process
  start       Start a process
  stop        Stop a process
```

### Flags
```bash
  -h, --help   help for routnd
```

### Use Cases
```bash
# Start a process
routnd start --name my-process --command "npm run dev"

# Start a process with a custom log file
routnd start --name my-process --command "npm run dev" --log-file /path/to/log/file.log

# Stop a process
routnd stop --name my-process

# List all running processes
routnd list

# Show logs for a process
routnd logs --name my-process
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

