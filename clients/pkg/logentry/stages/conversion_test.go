package conversion

import (
	"testing"
	"time"
)

func TestNewConversion(t *testing.T) {
	cfg := &ConversionConfig{}

	_, err := NewConversion(cfg)
	if err != nil {
		t.Errorf("NewConversion() failed with error: %v", err)
	}
}

func TestProcess(t *testing.T) {
	cfg := &ConversionConfig{}
	conversion, _ := NewConversion(cfg)

	labels := make(map[string]string)
	extracted := make(map[string]interface{})
	timestamp := time.Now()
	line := "This is a test log"

	conversion.Process(labels, extracted, &timestamp, &line)
	if line != "This is a converted test log" {
		t.Errorf("Process() failed to convert the log, expected: %s, got: %s", "This is a converted test log", line)
	}
}
