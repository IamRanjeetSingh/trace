package trace

import (
	"strings"
	"testing"
)

func testCallTrace() string {
	return Trace()
}

func Test_Trace_TraceDetailsNotEmpty(t *testing.T) {
	traceDetails := testCallTrace()
	if traceDetails == "" {
		t.Fatal("traceDetails are empty, expected non-empty details")
	}
}

func Test_Trace_TraceDetailsContainFuncName(t *testing.T) {
	traceDetails := testCallTrace()
	if !strings.Contains(traceDetails, "testCallTrace") {
		t.Fatal("traceDetails does not contain function name 'testCallTrace'")
	}
}

func Test_Trace_TraceDetailsContainFileName(t *testing.T) {
	traceDetails := testCallTrace()
	if !strings.Contains(traceDetails, "main_trace_test.go") {
		t.Fatal("traceDetails does not contain function name 'testCallTrace'")
	}
}
