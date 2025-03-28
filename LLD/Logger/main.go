package main

import (
	"fmt"
	"sync"

	dp "example.logger.com/differentPackage"
	"example.logger.com/logger"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			log := logger.GetLogger()
			log.WriteToLog(fmt.Sprintf("helloworld %d\n", i))
		}(i)
	}
	wg.Wait()
	tl := dp.TestLog{
		Log: logger.GetLogger(),
	}
	tl.PrintToTxt("putting an end to logs\n")
}
