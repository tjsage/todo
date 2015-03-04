package models

import (
	"testing"
)

func TestTaskNameMaxLength(t *testing.T) {
	//To Do: Don't trust that the Save method accurately tells us of it saved or not.
	task := &Task{Name: "ReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTaskReallyReallyReallyReallyReallyReallyLongTask"}
	if task.Save() != false {
		t.Fail()
	}
}
