package logger

import (
	"bytes"
	"log"
	"testing"
)

func TestInitSetsLoggers(t *testing.T) {
	Init()

	if Info == nil {
		t.Error("Info logger is not initialized")
	}
	if Warn == nil {
		t.Error("Warn logger is not initialized")
	}
	if Error == nil {
		t.Error("Error logger is not initialized")
	}
}

func TestLogOutput(t *testing.T) {
	var buf bytes.Buffer
	Info = newLogger(&buf, "INFO: ")

	Info.Println("This is a test log")

	output := buf.String()
	if output == "" {
		t.Error("Expected output from Info logger, got empty string")
	}
	if got, want := output[:6], "INFO: "; got != want {
		t.Errorf("Expected prefix %q, got %q", want, got)
	}
}

func newLogger(out *bytes.Buffer, prefix string) *log.Logger {
	return log.New(out, prefix, 0)
}
