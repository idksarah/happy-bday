package main

import (
	"os"

	"github.com/ansoni/termination"
	"github.com/nsf/termbox-go"
)

var letter = termination.Shape{
	"default": []string{`
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        | <-- envelope btw :<
|                      |
|______________________|
   press the up arrow!`,
		`
		
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        | <-- envelope btw :<
|                      |
|______________________|
   press the up arrow!`,
	},
}

var letter1 = termination.Shape{
	"default": []string{`
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        |
|                      |
|______________________|
 (type and press enter)`,
		`
		
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        |
|                      |
|______________________|
 (type and press enter)`,
	},
}

var letter2 = termination.Shape{
	"default": []string{`
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        |
|                      |
|______________________|
i'm happy you're here. i
 hope things are going 
         well`,
		`
		
 ______________________
|  \__            __/  |
|     \__      __/     |
|        \____/        |
|                      |
|______________________|
i'm happy you're here. i
 hope things are going 
         well`,
	},
}
var dots = termination.Shape{
	"default": []string{
		". . .",
		`
. . .`,
	},
}

var gotYouSomething = termination.Shape{
	"default": []string{
		`hey! i got you something
           :0`,
		`
hey! i got you something
           :0`,
	},
}

var gift = termination.Shape{
	"default": []string{`
		    _______________
  /----\\ /----\\ _/|
 _|-----\/-----||/| |
/_____/__/_____/| | |
|    |   |    | | | |
|    |   |    | | | |
|    |   |    | | |_/
|    |   |    | |_/
|____|___|____|_/  
     wow!!`, `

		    _______________
  /----\\ /----\\ _/|
 _|-----\/-----||/| |
/_____/__/_____/| | |
|    |   |    | | | |
|    |   |    | | | |
|    |   |    | | |_/
|    |   |    | |_/
|____|___|____|_/
     wow!!`,
	},
}

var yes1 = termination.Shape{
	"default": []string{
		`        open?
[X] yes    [] yes!!`,
		`
        open?
[X] yes    [] yes!!`,
	},
}

var yes2 = termination.Shape{
	"default": []string{
		`        open?
[] yes    [X] yes!!`,
		`
        open?
[] yes    [X] yes!!`},
}

var cake = termination.Shape{
	"default": []string{`
    *	      /\	         /\     - 
    	    ___||____/\____||___     |
        /   --    ||    --    \
    \	  |---------------------| 
        |---///---///---///---|
      *	|---------------------|   *
        |---///---///---///---|
    -	  |_____________________| \
      /      *    -      |

         happy (late) birthday 
         justin! i hope you're 
        having fun in san diego.
               i miss u!
  `, `
    *	      /\	         /\     \ 
    	    ___||____/\____||___     /
        /   --    ||    --    \
    |	  |---------------------| 
        |---///---///---///---|
      *	|---------------------|   *
        |---///---///---///---|
    \	  |_____________________| |
      -      *    \      /
	  
         happy (late) birthday 
         justin! i hope you're 
        having fun in san diego.
               i miss u!
  `,
		`
      *	   /\ 	        /\     | 
    		   ___||____/\____||___     -
    	   /   --    ||    --    \
     /  |---------------------| 
    	   |---///---///---///---|
    	 *	|---------------------|   *
    	   |---///---///---///---|
     |	 |_____________________| /
    	 \      *    |      -
	
    	   happy (late) birthday 
    	   justin! i hope you're 
    	  having fun in san diego.
    		        	i miss u!
`, `
       *     /\	         /\     / 
        		___||____/\____||___     \
        	/   --    ||    --    \
    -	   |---------------------| 
        	|---///---///---///---|
      *	 |---------------------|   *
        	|---///---///---///---|
    /	   |_____________________| -
      |      *    /      \
  
    	   happy (late) birthday 
    	   justin! i hope you're 
    	  having fun in san diego.
    		        i miss u!
`,
	},
}

type KirbyData struct {
	GotoX int
	GotoY int
}

func kirbyMovement(t *termination.Termination, e *termination.Entity, position termination.Position) termination.Position {
	return position
}

func printAt(input string, x int, y int) {
	termbox.Flush()
	print(input)
}

func main() {
	term := termination.New()
	term.FramesPerSecond = 2
	animation := term.NewEntity(termination.Position{50, 5, 0})
	animation.Shape = letter
	animation.Data = KirbyData{GotoX: 50, GotoY: 5}

	go term.Animate()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	var stage int = 0
	var stageMax int = 8
	userInput := "______!"
	// var name string

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				term.Close()
				os.Exit(0)
			}
			if ev.Key == termbox.KeyArrowUp {
				stage = min(stageMax, stage+1)
				term, animation = updateAnimation(stage, term, animation, 2)
			}
			if stage == 1 {
				if ev.Key == termbox.KeyEnter {
					termbox.Flush()
					// printAt("", 32, 0)
					// name = userInput
					termbox.SetCursor(0, 0)
					stage = min(stageMax, stage+1)
					term, animation = updateAnimation(stage, term, animation, 2)
					// print(name)
				} else {
					if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
						if len(userInput) > 0 {
							userInput = userInput[:len(userInput)-1]
						}
					} else {
						if userInput == "______!" {
							userInput = ""
						}
						userInput += string(ev.Ch)
					}
					printAt("hi "+userInput+"!", 32+len(userInput), 0)
				}
			}
		}
		// if ev.Key == termbox.KeyArrowDown {
		// 	stage = max(0, stage-1)
		// 	term, animation = updateAnimation(stage, term, animation, 1)
		// }
		// a way to go back would be a good idea but icl idc
	}
}

func updateAnimation(stage int, term *termination.Termination, animation *termination.Entity, framesPerSec int) (*termination.Termination, *termination.Entity) {
	animation.Die()
	animation = term.NewEntity(termination.Position{40, 5, 0})

	switch stage {
	case 0:
		animation.Shape = letter
	case 1:
		animation.Shape = letter1
		termbox.SetCursor(76, 12)
	case 2:
		animation.Shape = letter2
	case 3:
		animation.Shape = dots
	case 4:
		animation.Shape = gotYouSomething
	case 5:
		animation.Shape = gift
	case 6:
		animation.Shape = yes1
	case 7:
		animation.Shape = yes2
	case 8:
		animation.Shape = cake
	}

	term.FramesPerSecond = framesPerSec

	return term, animation
}
