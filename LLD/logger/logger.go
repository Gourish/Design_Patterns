package logger

import (
	"fmt"
	"os"
	"sync"
)

type Log struct {
	logFile *os.File
}

var log Log
var once sync.Once

func init() {
	log = Log{}
}

type ILog interface {
	WriteToLog(data string) error
}

func (log *Log) WriteToLog(data string) error {

	_, err := log.logFile.WriteString(data)
	return err

}

func GetLogger() ILog {
	once.Do(func() {
		logfile, err := os.Create("log.txt")
		if err != nil {
			fmt.Println("error creating log file ", err)
			return
		}
		log.logFile = logfile
	})
	return &log
}
