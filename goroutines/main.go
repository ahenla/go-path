package main

import (
	"fmt"
	"time"
)

func PrintMessage(text string, channel chan string) { // using the channel as an argument
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done" //sending data through the channel
}

func main() { // Main go routine or process.

	var Mych chan string // defining a channel with the chan keyword

	// go PrintMessage("Go is great", Mych)
	go PrintMessage("This is a go course", Mych)
	// go PrintMessage("gogogogogogogo!", Mych)
	response := <-Mych // waiting for the data from the channel.

	fmt.Println(response)
}

// Go routines are a way to rum multiple process inside the main
// process (main function). This feature can be accessed by any function
// putting the keyword go before its call.
// If the main process is finished, any go routine will be killed with it.
//
//
// Channels are a special variable that can be defined using the chan
// keyword. Channels are used to send data between go routines. So a
// thread will wait for the channel to bring some data before continue
// the execution. This can be used as a way to keep the main go routine
// open.
//Channels can be buffered too. Which means that is possible to send
//multiple pieces of data through the same channel and they will be
//avalilable to be accessed when reading from the channel (the last
//piece won't replace the previous). In order to do this ypu can specify
//the number of pieces of data you want to send at a time as a second
//argument in the channel declaration.
