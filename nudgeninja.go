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
	flag.IntVar(&seconds, "sec", 300, "An integer representing seconds")

	// Define a custom help flag
	showHelp := flag.Bool("help", false, "Show help message")
	flag.Parse()

	// Check if the help flag was provided, and if so, display the help message
	if *showHelp {
		fmt.Println("Usage of the program:")
		fmt.Println("  -sec int")
		fmt.Println("        An integer representing seconds. Default: 300")
		fmt.Println("  -help")
		fmt.Println("        Show help message")
		os.Exit(0)
	}

	timer(seconds)
}

func timer(seconds int) {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	var quit = make(chan bool)

	x, y := robotgo.Location()

	go func() {
		for {
			select {
			case <-ticker.C:
				x, y = move(x, y)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for the quit channel
	<-quit
}

// Moves the mouse to a random location on the screen if the mouse hasn't moved since the last call
func move(prevX int, prevY int) (int, int) {
	currentX, currentY := robotgo.Location()

	if currentX == prevX && currentY == prevY {
		screenX, screenY := robotgo.GetScreenSize()
		randX, randY := rand.Intn(screenX), rand.Intn(screenY)
		robotgo.Move(randX, randY)
		return randX, randY
	}

	return currentX, currentY
}
