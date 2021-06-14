package main

import (
	"fmt"
	"time"
	"math/rand"
    "math"
    "os"

	"github.com/go-vgo/robotgo"
	// hook "github.com/robotn/gohook"
)

var avgTime float64
var avgTimeCount int
var otherWindowX, otherWindowY int
var mainX, mainY int

func main() {
	avgTimeCount = 0
	fmt.Print("How many items: ")
	var iterations int
    fmt.Scanln(&iterations)
	fmt.Println("\nMove mouse to item/spell location")
	delay(4)
	mainX, mainY = robotgo.GetMousePos()
	fmt.Println("pos:", mainX, mainY)
	fmt.Println("\nMove mouse to second window location")
	delay(4)
	otherWindowX, otherWindowY = robotgo.GetMousePos()
	fmt.Println("pos:", otherWindowX, otherWindowY)
	autoClick(iterations)
}

func autoClick(iterations int) {
	// fmt.Println("Key press listener starting...")
	// evChan := hook.Start()
	// defer hook.End()

	count := 1
	avgTime = 0
	avgTimeCount = 0
	if (iterations <= 1) {
		os.Exit(iterations)
	}

	randomIterations := iterations
	if (iterations > 150) {
	    randomIterations = rand.Intn(int(math.Round(float64(iterations/2))))+rand.Intn(iterations)
		if (randomIterations > iterations) {
			randomIterations = iterations
		}
	}

	pauseInterval := rand.Intn(int(math.Round(float64(randomIterations/2))))+11
    fmt.Printf("Pause at: ", pauseInterval)

	robotgo.MoveMouse(mainX+rand.Intn(120)+71, mainY+rand.Intn(120)+167)
	delay(0)
	mouseMoveRandom()
	// clickSpellBook(mainX, mainY)

	avgTime = 0
	// for ev := range evChan {
	for count <= randomIterations {
		// if ev.Rawcode == 32 { //Spacebar is 32
		// 	fmt.Println("space bar")
		// } else {
		timeLeft := float64(randomIterations-count+3)*(5.8+0.2)
		fmt.Print("=========== Reamining Time: ~")
		if (timeLeft > 60.0) {
			timeLeft = timeLeft/60
			fmt.Println(math.Ceil(timeLeft), " min ==============")
			math.Ceil(timeLeft)
		} else {
			fmt.Println(fmt.Sprintf("%.2f", timeLeft), " seconds ==============")
		}

		fmt.Println("Action: ", count, " of ", randomIterations, " -- Pause at ", pauseInterval)
		if (rand.Intn(10) > 9) {
			mouseMoveRandom()
		}

		clickTwice()
		if (count == pauseInterval && count < randomIterations) {
			fmt.Println("\n======= pausing =========")
			if (pauseInterval < randomIterations) {
				pauseInterval = ((rand.Intn(7)*2)+count+10)
				fmt.Println("Next pause at: ", pauseInterval)
			}

			messUp("click")
			clickOtherWindow()

			fmt.Println("======= paused ==========\n")
		}

		count++

		// fmt.Println("--- Pause now by pressing Space Bar ---")
		// hook.Register(hook.KeyDown, []string{" "}, func(e hook.Event) {
		// 	fmt.Println("Releasing for 18 - 20 seconds")
		// 	time.Sleep(10 * time.Second)
		// 	fmt.Println("~10 seconds left")
		// 	time.Sleep(8 * time.Second)
		// 	fmt.Println("~3 seconds left")
		// 	delay(3)
		// 	fmt.Println("Starting again")
		// 	clickOtherWindow()
		// 	hook.End()
		// })

		// s := hook.Start()
		// <-hook.Process(s)
	}

	totalTimes(randomIterations)
	fmt.Print("remaining iterations ", iterations-randomIterations, " | New Iterations (0 to continue):")
	var newIterations int
    fmt.Scanln(&newIterations)
    if (newIterations <= 0) {
		newIterations = iterations-randomIterations
    }

    autoClick(newIterations)
}

func clickTwice() {
	delay(1)
	robotgo.Click()
	messUp("both")
	fmt.Print("Clicking 1 -- ")
	delay(1)
	robotgo.Click()
	messUp("both")
	fmt.Print("Clicking 2 -- ")
	delay(1)
}

func delay(delayTime float64) {
    rand.Seed(time.Now().UnixNano())
	n := (rand.Float64()+0.7) // n will be between 0 and 10
	if (delayTime >= 3) {
		n=n+delayTime
	}

    if (n < 1.2) {
	    n = n+delayTime
     	if (delayTime == 0) {
	    	n = n+0.5
	    }
    }

	// n=n+rand.Float64()+1

	if (delayTime < 3.0) {
	    avgTime = avgTime+n
	    avgTimeCount++
	}

	fmt.Println("Sleeping for seconds:", n)
	time.Sleep(time.Duration(n)*time.Second)
	messUp("time")
}

func clickSpellBook(mainX int, mainY int) {
	robotgo.MoveMouseSmooth(mainX+181, mainY+231, (rand.Float64()+1)*3.1, (rand.Float64()+1)*3.3)
	robotgo.Click()
	delay(1)
	robotgo.MoveMouseSmooth(mainX+179, mainY+189, (rand.Float64()+1)*3.2, (rand.Float64()+1)*3.4)
	robotgo.Click()
	delay(1)
}

func mouseMoveRandom() {
	x := ((rand.Intn(4)*-1 - rand.Intn(4)*-1))+mainX
	y := ((rand.Intn(3)*-1 - rand.Intn(3)*-1))+mainY
	robotgo.MoveMouseSmooth(x, y, (rand.Float64()+1)*2.9, (rand.Float64()+1)*1.6)
}

func clickOtherWindow() {
	robotgo.MoveMouseSmooth(rand.Intn(6000), rand.Intn(1800), (rand.Float64()), (rand.Float64()))
	robotgo.MoveMouseSmooth(otherWindowX, otherWindowY, (rand.Float64())*1.7, (rand.Float64())*2.1)
	robotgo.Click()
	delay(5)
	robotgo.MoveMouseSmooth(rand.Intn(6000), rand.Intn(1800), (rand.Float64()), (rand.Float64()))
	x := ((rand.Intn(4)*-1 - rand.Intn(4)*-1))+mainX
	y := ((rand.Intn(3)*-1 - rand.Intn(3)*-1))+mainY
	robotgo.MoveMouseSmooth(x, y, (rand.Float64()+1)*1.3, (rand.Float64()+1)*1.1)
	delay(1)
	mouseMoveRandom()
}

func messUp(messUpType string) {
	switch messUpType {
		case "click":
			if (rand.Intn(100) >= 80) {
				n := rand.Float64()
				fmt.Print("Clicking 1A -- ")
				time.Sleep(time.Duration(n)*time.Second)
				fmt.Println("Bonus Sleeping for seconds:", n)
				robotgo.Click()
			}
		case "time":
			if (rand.Intn(100) >= rand.Intn(20)+80) {
				pauseX, pauseY := robotgo.GetMousePos()
				fmt.Print("Bonus ")
				delay(2)
				if (rand.Intn(100) >= 95) {
					if (rand.Intn(100) >= 50) {
						clickOtherWindow()
					}

					fmt.Print("Super Bonus ")
					delay(6)
					if (rand.Intn(100) >= 99) {
						clickOtherWindow()
						fmt.Print("Super Mega Bonus ")
						delay(12)
						if (rand.Intn(1000) >= 999) {
							fmt.Print("Super Mega Supreme Bonus ")
							delay(30)
						}
					}
				}

				robotgo.MoveMouseSmooth(pauseX, pauseY, (rand.Float64())*1.3, (rand.Float64())*1.1)
				robotgo.Click()
			}
		case "both":
			if (rand.Intn(100) >= 80) {
				n := rand.Float64()+0.4
				fmt.Print("Clicking 1A -- ")
				time.Sleep(time.Duration(n)*time.Second)
				fmt.Println("Bonus Sleeping for seconds:", n)
				robotgo.Click()
				if (rand.Intn(100) >= 85) {
					n := rand.Float64()+0.4
					time.Sleep(time.Duration(n)*time.Second)
					fmt.Print("Clicking 1B -- ")
					robotgo.Click()
					fmt.Println("Super Bonus Sleeping for seconds:", n)
				}
			}
	}
}

func totalTimes(iterations int) {
	fmt.Println("\n\n=========================")
	avgTimeTime := avgTime/float64(avgTimeCount)
	fmt.Println("finished all iterations. Average time for each action: ", avgTimeTime)
	avgTimeAction := avgTime/float64(iterations)
	fmt.Println("finished all iterations. Average time per action: ", avgTimeAction)
	fmt.Println("=========================\n\n")
}
