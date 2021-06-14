package main

import (
	"fmt"
	"time"
	"math/rand"
    // "math"
    "os"

	"github.com/go-vgo/robotgo"
	// hook "github.com/robotn/gohook"
)

var avgTime float64
var avgTimeCount int
var otherMouseX, otherMouseY int
var mainX, mainY int

func main() {
	avgTimeCount = 0
	n := 3
	fmt.Print("How many items: ")
	var iterations int
    fmt.Scanln(&iterations)
	fmt.Println("\nMove mouse to item/spell location")
	time.Sleep(time.Duration(n)*time.Second)
	mainX, mainY = robotgo.GetMousePos()
	fmt.Println("pos:", mainX, mainY)
	fmt.Println("\nMove mouse to second window location")
	time.Sleep(time.Duration(n)*time.Second)
	otherMouseX, otherMouseY = robotgo.GetMousePos()
	fmt.Println("pos:", otherMouseX, otherMouseY)

	autoClick(iterations)	
}

func autoClick(iterations int) {
	count := 1
	avgTime = 0
	avgTimeCount = 0
	if (iterations <= 1) {
		os.Exit(iterations)
	}

	robotgo.MoveMouse(mainX, mainY)
	for count <= iterations {

		robotgo.MoveMouse(mainX, mainY)
		fmt.Println("Action: ", count, " of ", iterations)
		mouseMoveRandom()
		count++
	}
}

func mouseMoveRandom() {
	// foo := 0
	count := 1
	max := 25
	currentX, currentY := robotgo.GetMousePos()
	differenceX := (currentX-otherMouseX)/max
	differenceY := (currentY-otherMouseY)/max
	if (currentX < otherMouseX) {
		differenceX = (otherMouseX-currentX)/max
		// mouseX = currentX-(differenceX)
	}

	if (currentY < otherMouseY) {
		differenceY = (otherMouseY-currentY)/max
		// mouseY = currentY-(differenceY)
	}
	
	for count <= max {
		// max = max - count
		mouseX := currentX+(differenceX*count)
		mouseY := currentY+(differenceY*count)
	
		if (differenceX < 0 || differenceY < 0) {
			os.Exit(-1)
		}
	
		fmt.Println(differenceX+rand.Intn(4), differenceY, mouseX, mouseY)
		count++
		if (rand.Intn(10) <= 5) {
			mouseX = mouseX+rand.Intn(10)
		} else {
			mouseX = mouseX-rand.Intn(10)
		}

		if (rand.Intn(10) <= 5) {
			mouseY = mouseY+rand.Intn(10)
		} else {
			mouseY = mouseY-rand.Intn(10)
		}

		robotgo.MoveMouseSmooth(mouseX, mouseY, 0.1, 0.1)
	}

	// x := ((rand.Intn(4)*-1 - rand.Intn(4)*-1))+mainX
	// y := ((rand.Intn(3)*-1 - rand.Intn(3)*-1))+mainY
	// robotgo.MoveMouse(x, y)
	// x2 := ((rand.Intn(4)*-1 - rand.Intn(4)*-1))+otherMouseX
	// y2 := ((rand.Intn(3)*-1 - rand.Intn(3)*-1))+otherMouseY
	// difference := 0
	// differenceX := (x - x2)/10
	// differenceY := (y - y2)/10
	// if (y>=y2) {
	// 	differenceX = (x - x2)/10
	// 	differenceY = (y - y2)/10
	// } else if (y<y2) {
	// 	differenceX = (x2 - x)/10
	// 	differenceY = (y2 - y)/10
	// }

	// for count<=10 {
	// 	count++
	// 	x2, foo = robotgo.GetMousePos()
	// 	if (rand.Intn(9)+1 >= 5) {
	// 		foo = (foo+differenceY)+(10+rand.Intn(10))
	// 	} else {
	// 		foo = (foo+differenceY)-(10+rand.Intn(10))
	// 	}

	// 	robotgo.MoveMouseSmooth(differenceX+(x2 - x)/10, foo, 0.1, 0.1)
	// }

	// robotgo.MoveMouseSmooth(x2, y2, 0.1, 0.1)
}

func messUp(messUpType string) {
	switch messUpType {
		case "click":
			if (rand.Intn(100) >= 80) {
				n := rand.Float64()		
				fmt.Print("Clicking 1A -- ")
				time.Sleep(time.Duration(n)*time.Second)
				fmt.Println("Bonus Sleeping for seconds:", n)
				//// robotgo.Click()
			}
		case "time":
			if (rand.Intn(100) >= rand.Intn(20)+80) {
				pauseX, pauseY := robotgo.GetMousePos()
				fmt.Print("Bonus ")
				// delay(2)
				if (rand.Intn(100) >= 80) {
					fmt.Print("Super Bonus ")
					// delay(6)
					if (rand.Intn(100) >= 99) {
						fmt.Print("Super Mega Bonus ")
						// delay(12)
						if (rand.Intn(1000) >= 999) {
							fmt.Print("Super Mega Supreme Bonus ")
							// delay(30)
						}
					}
				}

				robotgo.MoveMouseSmooth(pauseX, pauseY, (rand.Float64())*1.3, (rand.Float64())*1.1)
				//// robotgo.Click()
			}
		case "both":
			if (rand.Intn(100) >= 80) {
				n := rand.Float64()+0.4		
				fmt.Print("Clicking 1A -- ")
				time.Sleep(time.Duration(n)*time.Second)
				fmt.Println("Bonus Sleeping for seconds:", n)
				//// robotgo.Click()
				if (rand.Intn(100) >= 85) {
					n := rand.Float64()+0.4
					time.Sleep(time.Duration(n)*time.Second)
					fmt.Print("Clicking 1B -- ")
					//// robotgo.Click()
					fmt.Println("Super Bonus Sleeping for seconds:", n)
				}
			}
	}
}