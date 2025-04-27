package main

import (
	"fmt"
	"os"

	"github.com/ansoni/termination"
	"github.com/nsf/termbox-go"
)

var letter = termination.Shape{
	"default": []string{
		` ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        |
|                      |
|______________________|`,
		`  ______________________
 |  \__            __/  |
 |     \__      __/     |
 |        \____/        |
 |                      |
 |______________________|`,
	},
}

var dots = termination.Shape{
	"default": []string{
		". . .",
		" . . .",
	},
}

var gotYouSomething = termination.Shape{
	"default": []string{
		"hey! i got you something",
		" hey! i got you something",
	},
}

var gift = termination.Shape{
	"default": []string{},
}

var yesOrYes = termination.Shape{
	"default": []string{},
}

var cake = termination.Shape{
	"default": []string{},
}

type KirbyData struct {
	GotoX int
	GotoY int
}

func kirbyMovement(t *termination.Termination, e *termination.Entity, position termination.Position) termination.Position {
	return position
}

func printAt(input interface{}) {
	fmt.Print("\r\033[K")
	str := fmt.Sprintf("%v", input)
	termbox.SetCursor(0, 0)
	print(str)
}

func main() {
	// Initialize the termination object and the animation entity
	term := termination.New()
	term.FramesPerSecond = 2
	animation := term.NewEntity(termination.Position{40, 5, 0})
	animation.Shape = letter
	animation.Data = KirbyData{GotoX: 40, GotoY: 5}

	// Start the animation in the background (only once)
	go term.Animate()
	printAt("hi!!")

	// Setup termbox input mode
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	var stage int = 0

	// Main input loop
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				term.Close()
				os.Exit(0)
			}
			if ev.Key == termbox.KeyArrowUp {
				stage++
				term, animation = updateAnimation(stage, term, animation, 2)
			}
			// if ev.Key == termbox.KeyArrowDown {
			// 	stage = max(0, stage-1)
			// 	term, animation = updateAnimation(stage, term, animation, 1)
			// }
			// a way to go back would be a good idea but icl idc
		}
	}
}

// updateAnimation updates the shape of the animation based on the stage
func updateAnimation(stage int, term *termination.Termination, animation *termination.Entity, framesPerSec int) (*termination.Termination, *termination.Entity) {
	// Remove the old animation
	animation.Die() // This cleans up the old animation entity

	// Create a new animation entity
	animation = term.NewEntity(termination.Position{40, 5, 0})

	// Update the shape based on the stage
	switch stage {
	case 0:
		animation.Shape = letter
	case 1:
		animation.Shape = dots
	}

	// Set the frames per second for the animation
	term.FramesPerSecond = framesPerSec

	// No need to call go term.Animate() here, it has already been started

	return term, animation
}

// max is a helper function to prevent stage from going below 0
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
