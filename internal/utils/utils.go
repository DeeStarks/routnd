package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func SplitCommand(command string) (string, []string) {
	// Split the string command into: command, args
	split := strings.Split(command, " ")
	return split[0], split[1:]
}

func ViewFile(fileName string, tail int, follow bool) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string

	if tail > 0 {
		// Get the file size
		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}
		fileSize := fileInfo.Size()

		// Get the tail
		buffer := make([]byte, fileSize)
		file.Read(buffer)
		lines = strings.Split(string(buffer), "\n")
		start := len(lines) - tail
		if start < 0 {
			start = 0
		}
		end := len(lines) - 1 // -1 to remove the last empty line
		if end < 0 {
			end = 0
		}
		lines = lines[start:end]
	} else {
		// Get all the lines
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	// Print the lines
	for _, line := range lines {
		fmt.Println(line)
	}

	// Follow the logs
	if follow {
		// Seek to the end of the file
		file.Seek(0, io.SeekEnd)

		// Read the file
		reader := bufio.NewReader(file)
		for {
			// Read a line
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					time.Sleep(1 * time.Second)
					continue
				}
				panic(err)
			}

			// Print the line
			fmt.Print(line)
		}
	}
}

func GetCurrentTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
