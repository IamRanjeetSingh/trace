package trace

import (
	"fmt"
	"path"
	"runtime"
)

func Trace() (traceDetails string) {
	programCounter, fullFilepath, line, ok := runtime.Caller(1)
	if !ok {
		panic("Error while getting runtime caller")
	}
	_, filename := path.Split(fullFilepath)
	function := runtime.FuncForPC(programCounter)
	funcName := function.Name()
	return fmt.Sprintf("%s:%d - %s", filename, line, funcName)
}
