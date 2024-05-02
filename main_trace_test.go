package trace

import (
	"strings"
	"testing"
)

func testCallTrace() (traceDetails string, ok bool) {
	return Trace()
}

func Test_Trace_OkIsTrue(t *testing.T) {
	_, ok := testCallTrace()
	if !ok {
		t.Fatal("ok was false, expected true")
	}
}

func Test_Trace_TraceDetailsNotEmpty(t *testing.T) {
	traceDetails, _ := testCallTrace()
	if traceDetails == "" {
		t.Fatal("traceDetails are empty, expected non-empty details")
	}
}

func Test_Trace_TraceDetailsContainFuncName(t *testing.T) {
	traceDetails, _ := testCallTrace()
	if !strings.Contains(traceDetails, "testCallTrace") {
		t.Fatal("traceDetails does not contain function name 'testCallTrace'")
	}
}

func Test_Trace_TraceDetailsContainFileName(t *testing.T) {
	traceDetails, _ := testCallTrace()
	if !strings.Contains(traceDetails, "main_trace_test.go") {
		t.Fatal("traceDetails does not contain function name 'testCallTrace'")
	}
}
