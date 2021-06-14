package main

import (
	"fmt"
	"time"
    "bufio"
    "log"
    "os"
    "strconv"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("hook start...")
	evChan := hook.Start()
	defer hook.End()
    
    file, err := os.Create("./temp.txt")
    if err != nil {
        log.Fatal(err)
    }

	var writer = bufio.NewWriter(file)
    // writer = bufio.NewWriterSize(writer, 50000) // couldnt get this to work

    runs := 1
	avgTime := time.Now().Unix()
    writer.WriteString("\n")
	writer.Flush()

	for ev := range evChan {
		if ((ev.Kind) == 9) {
       		writer.WriteString(strconv.Itoa((int(ev.X)*-1))) // having it as *-1 because monitors work in 4th quad
       		writer.WriteString(", ")
       		writer.WriteString(strconv.Itoa(int(ev.Y)))
       		writer.WriteString("\n")
			writer.Flush()
		} else  if ((ev.Kind) == 7) {
	        fmt.Printf(strconv.Itoa(int(ev.Kind)))
			t := time.Now().Unix()
			runs++
			avgTime = t+avgTime
	        fmt.Println("time %d", t)
	        fmt.Println(strconv.Itoa(runs), ": %d", (float64(avgTime)/float64(runs))/float64(1000000000))
		}
	}
}

