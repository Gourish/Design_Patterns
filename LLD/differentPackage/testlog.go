package differentpackage

import "example.logger.com/logger"

type TestLog struct {
	Log logger.ILog
}

func (tl *TestLog) PrintToTxt(str string) {
	tl.Log.WriteToLog(str)
}
