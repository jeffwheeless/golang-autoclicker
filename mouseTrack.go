package main

import (
	"fmt"
	"time"
    // "bufio"
	"math/rand"
    // "log"
    // "os"
    "strconv"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func get() {
	time.Sleep(2 * time.Second)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

func main() {
    // file, err := os.Create("./temp.txt")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // writer := bufio.NewWriter(file)
    // linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    // for _, line := range linesToWrite {
    // }
    // writer.Flush()

	fmt.Println("hook start...")
	evChan := hook.Start()
	defer hook.End()

	// count := 0
	// max := 20

	// file, err := os.Create("./temp.txt")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // writer := bufio.NewWriter(file)

    // // writer = bufio.NewWriterSize(writer, 50000)
    // linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    // for _, line := range linesToWrite {
    //     bytesWritten, err := writer.WriteString(line + "\n")
    //     if err != nil {
    //         log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
    //     }
    //     fmt.Printf("Bytes Written: %d\n", bytesWritten)
    //     fmt.Printf("Available: %d\n", writer.Available())
    //     fmt.Printf("Buffered : %d\n", writer.Buffered())
    // }
    // writer.Flush()

	// // file, err := os.Create("./temp2.txt")
 // //    if err != nil {
 // //        log.Fatal(err)
 // //    }
 // //    writer := bufio.NewWriter(file)
 //    // for _, line := range linesToWrite {
	// for ev := range evChan {
	// 	fmt.Println("hook: ", ev)
	// 	fmt.Println("kind: ", strconv.Itoa(int(ev.Kind)))
	// 	if ((ev.Kind) == 9 || (ev.Kind) == 1) {
	// 	// if (runs/3 == 0) {
 //       		bytesWritten, err := writer.WriteString(strconv.Itoa(int(ev.X)))
 //       		bytesWritten, err = writer.WriteString(", ")
 //       		bytesWritten, err = writer.WriteString(strconv.Itoa(int(ev.Y)))
 //       		bytesWritten, err = writer.WriteString("\n")
 //       		// bytesWritten, err := writer.WriteString("1")
	//         if err != nil {
	//             log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	//         }
	// 		writer.Flush()
	//         // fmt.Printf(strconv.Itoa(int(ev.Kind)))
	//         // fmt.Printf(strconv.Itoa(int(ev.X)))
	//         // fmt.Printf(strconv.Itoa(int(ev.Y)))
	//         fmt.Printf("Bytes Written: %d\n", bytesWritten)
	//         fmt.Printf("Available: %d\n", writer.Available())
	//         fmt.Printf("Buffered : %d\n", writer.Buffered())
	// 	}  else {
	// 		writer.Flush()
	// 	}

 //        // bytesWritten, err := writer.WriteString(line + "\n")
 //        // if err != nil {
 //        //     log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
 //        // }
 //        // fmt.Printf("Bytes Written: %d\n", bytesWritten)
 //        // fmt.Printf("Available: %d\n", writer.Available())
 //        // fmt.Printf("Buffered : %d\n", writer.Buffered())
 //    }




    file, err := os.Create("./temp.txt")
    if err != nil {
        log.Fatal(err)
    }

	var writer = bufio.NewWriter(file)
    // writer = bufio.NewWriterSize(writer, 50000)

    runs := 1
	avgTime := time.Now().Unix()
    bytesWritten, err := writer.WriteString("\n")
	writer.Flush()

	for ev := range evChan {
        // fmt.Printf(strconv.Itoa(int(ev.Y)))
		    // var writer = bufio.NewWriter(file)
		if ((ev.Kind) == 9) {
       		writer.WriteString(strconv.Itoa((int(ev.X)*-1)))
       		writer.WriteString(", ")
       		writer.WriteString(strconv.Itoa(int(ev.Y)))
       		writer.WriteString("\n")
	        // fmt.Printf(strconv.Itoa(int(ev.Kind)))
	        // fmt.Printf(strconv.Itoa(int(ev.X)))
	        // fmt.Printf(strconv.Itoa(int(ev.Y)))
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

