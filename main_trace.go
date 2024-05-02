package trace

import (
	"fmt"
	"path"
	"runtime"
)

func Trace() (traceDetails string, ok bool) {
	programCounter, fullFilepath, line, ok := runtime.Caller(1)
	if !ok {
		return "", false
	}
	_, filename := path.Split(fullFilepath)
	function := runtime.FuncForPC(programCounter)
	funcName := function.Name()
	return fmt.Sprintf("%s:%d - %s", filename, line, funcName), true
}
