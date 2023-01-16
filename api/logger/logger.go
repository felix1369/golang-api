package logger

import (
	"log"
	"os"
)

// Type is variable for log
type Type struct {
	DefaultLog *log.Logger
	StompLog   *log.Logger
}

// LogType for usage
var LogType = Type{}

func init() {
	// set location of log file
	file, err := os.OpenFile("log_file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	LogType.DefaultLog = log.New(file, "", log.LstdFlags|log.Lshortfile)
	stompLogfile, err := os.OpenFile("stomp_log_file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	LogType.StompLog = log.New(stompLogfile, "", log.LstdFlags|log.Lshortfile)
}
