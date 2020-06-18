package main

import (
	"log"
	"os"
)

var debugLog *log.Logger

func main() {

	debug := false
	if debug {
		debugLog = log.New(os.Stdout, "[Debug]", log.Llongfile)
	} else {
		logFile, err := os.Create("/dev/null")
		if err != nil {
			log.Fatalln("open file error !")
		}
		debugLog = log.New(logFile, "[Debug]", log.Llongfile)
		defer logFile.Close()
	}

	debugLog.Println("A debug message here")

}
