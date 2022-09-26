package main

import (
	"math"

	"github.com/zyellowhorse/go-template/gamercade"
)

var (
	frameCounter uint  = 0
	xPos         int32 = 0
	yPos         int32 = 0
)

// Called once, at the start of the game.
func Init() {
	gamercade.ConsoleLog("Hello, from Golang!")

	xPos = int32(gamercade.Width() / 2)
	yPos = int32(gamercade.Height() / 2)
}

// Called once every frame, before draw.
func Update() {
	// Print a message if the user presses the A button.
	// This defaults to the U key on the keyboard.
	if gamercade.ButtonAPressed(0) {
		state := gamercade.ButtonAPressed(0)
		if state {
			gamercade.ConsoleLog("Pressed A.")
		}
	}

	// Let's move the pixel with the arrow keys
	// Handle up/down motion
	if gamercade.ButtonUpHeld(0) {
		state := gamercade.ButtonUpHeld(0)
		if state {
			yPos -= 1
		}
	}

	if gamercade.ButtonDownHeld(0) {
		state := gamercade.ButtonDownHeld(0)
		if state {
			yPos += 1
		}
	}

	// And repeat for left/right
	if gamercade.ButtonLeftHeld(0) {
		state := gamercade.ButtonLeftHeld(0)
		if state {
			xPos -= 1
		}
	}

	if gamercade.ButtonRightHeld(0) {
		state := gamercade.ButtonRightHeld(0)
		if state {
			xPos += 1
		}
	}

	// Update the frame counter to keep the animation looping
	frameCounter += 1
}

// Called once every frame, after update.
func Draw() {
	// Clear screen function takes a GraphicsParameters as a parameter,
	// so let's make one.
	clearColor := gamercade.ColorIndex(0)

	// Now, we can clear the screen.
	gamercade.ClearScreen(clearColor)

	// Let's draw a pixel.
	pixelColor := gamercade.ColorIndex(16)
	gamercade.SetPixel(pixelColor, xPos, yPos)

	// Let's draw a spinning pixel.
	spinningPixelColor := gamercade.ColorIndex(9)

	// Make it spin around
	frame := float64(frameCounter)
	x := math.Sin(frame*0.1) * 25.0
	y := math.Cos(frame*0.1) * 25.0

	x += float64(xPos)
	y += float64(yPos)

	// Draw the spinning pixel
	gamercade.SetPixel(spinningPixelColor, int32(x), int32(y))
}
