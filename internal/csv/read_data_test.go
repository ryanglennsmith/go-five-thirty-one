package csv

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	filePath := "../../.csv/test-data.csv"
	fmt.Println(filePath)
	data, err := ReadData(filePath)
	if err != nil {
		t.Errorf("Error reading data: %s", err)
	}
	if len(data) != 48 {
		t.Errorf("Expected 48 records, but got %d", len(data))
	}

	// Test the first record
	firstRecord := data[0]
	if firstRecord.Lift != "DL" {
		t.Errorf("Expected lift to be DL, but got %s", firstRecord.Lift)
	}
	if firstRecord.WeekSet != "W1S1" {
		t.Errorf("Expected weekset to be W1S1, but got %s", firstRecord.WeekSet)
	}
	if firstRecord.Weight != 120 {
		t.Errorf("Expected weight to be 120, but got %f", firstRecord.Weight)
	}
}