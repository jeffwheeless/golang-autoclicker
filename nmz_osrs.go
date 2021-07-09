package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	// hook "github.com/robotn/gohook"
)

var mainX, mainY int
var timeTotal, avgTime, shortestTime, longestTime float64
var dryRun bool

func main() {
	runs := 1500.0
	shortestTime = runs
	longestTime = 0
	var placeHolder int
	fmt.Println("\nPosition 1 then enter")
	fmt.Scanln(&placeHolder)
	mainX, mainY = robotgo.GetMousePos()
	fmt.Println("pos:", mainX, mainY)
	dryRun = true
	avgTime = autoClick(runs)
	fmt.Println("average time: ", avgTime)
	fmt.Print("\nHow long do you want the thing to run (in min):")
	fmt.Scanln(&runs)
	fmt.Println("Total runs: ", runs*(avgTime/60))
	dryRun = false
	avgTime = autoClick(runs)
	fmt.Println("average time: ", avgTime)
}

func autoClick(runs float64) float64 {
	x := 0.0
	timeTotal = 0.0
	for x <= runs {
		x++
		if dryRun == false {
			if x/5 == 0 {
				fmt.Println("shortest recorded thus far:", shortestTime)
				fmt.Println("longest recorded thus far:", longestTime)
			}

			fmt.Println("\n\nrun: ", x, " of ", runs)
			robotgo.MoveMouseSmooth(mainX+rand.Intn(5), mainY+rand.Intn(5), (rand.Float64() + 1), (rand.Float64() + 1))
		}

		delay(1.8 - (rand.Float64() * .8))
		if dryRun == false {
			robotgo.Click()
		}

		delay(1.7 - (rand.Float64() * .6))
		if dryRun == false {
			robotgo.MoveMouseSmooth(mainX+rand.Intn(5), mainY+rand.Intn(5), (rand.Float64()+1)*3.1, (rand.Float64()+1)*3.3)
			robotgo.Click()
		}

		delay(rand.Float64() + float64(rand.Intn(10)+rand.Intn(5)+40))
	}

	return (timeTotal / float64(runs))
}

func delay(sleepTime float64) {
	// sleepTime = sleepTime + 0.02
	if sleepTime < 1.02 {
		sleepTime = 1.03 - (rand.Float64() / 100)
	}

	timeTotal += sleepTime
	if shortestTime > sleepTime {
		shortestTime = sleepTime
		fmt.Println("New shortest time: ", shortestTime)
	}

	if longestTime < sleepTime {
		longestTime = sleepTime
		fmt.Println("New longest time: ", longestTime)
	}

	if dryRun == false {
		fmt.Println("Sleep for: ", sleepTime)
		if sleepTime > 30 {
			time.Sleep(time.Duration(sleepTime-3.0) * time.Second)
			warn := 3
			fmt.Print("\t")
			for warn > 0 {
				fmt.Print(" ...", warn)
				warn--
				time.Sleep(time.Duration(float64(1)) * time.Second)
			}
			fmt.Print("\n")
		} else {
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}
}
