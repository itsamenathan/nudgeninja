package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Define a command line flag called 'sec' to capture an integer value
	var seconds int
	flag.IntVar(&seconds, "sec", 60, "An integer representing seconds")

	// Define a custom help flag
	showHelp := flag.Bool("help", false, "Show help message")

	// Parse the command line flags
	flag.Parse()

	// Check if the help flag was provided, and if so, display the help message
	if *showHelp {
		fmt.Println("Usage of the program:")
		fmt.Println("  -sec int")
		fmt.Println("        An integer representing seconds. Default: 60")
		fmt.Println("  -help")
		fmt.Println("        Show help message")

		// Exit the program after displaying the help message
		os.Exit(0)
	}

	// Run our move function
	move(seconds)
}

func move(seconds int) {
	robotgo.MouseSleep = 100

	for {
		// Get start location
		start_x, start_y := robotgo.Location()

		time.Sleep(time.Duration(seconds) * time.Second)

		// Get our current location
		current_x, current_y := robotgo.Location()

		// If the start and current locations are the same, move the mouse
		if start_x == current_x && start_y == current_y {
			robotgo.Move(rand.Intn(1000), rand.Intn(1000))
		}
	}
}
