package cmd_test

import (
	. "cmd"
	"testing"
)

func TestRunTrue(t *testing.T) {
	cmd := Command("true")
	err := cmd.Run()
	if err != nil {
		t.Error("Run error: ", err)
	}
	pid, err := cmd.Pid()
	if err != nil {
		t.Error("Pid error: ", err)
	}
	if pid <= 0 {
		t.Error("Pid Expected >0, got ", pid)

	}
	ex, err := cmd.ExitStatus()
	if err != nil {
		t.Error("ExitStatus error: ", err)
	}
	if ex != 0 {
		t.Error("ExitStatus Expected 0, got ", ex)
	}
	_, err = cmd.Signal()
	if err != ProcessNotSignaled {
		t.Error("Signal Expected ProcessNotSignaled, got ", err)
	}

}

func TestRunFalse(t *testing.T) {
	cmd := Command("false")
	err := cmd.Run()
	if err != nil {
		t.Error("Run error: ", err)
	}
	pid, err := cmd.Pid()
	if err != nil {
		t.Error("Pid error: ", err)
	}
	if pid <= 0 {
		t.Error("Pid Expected >0, got ", pid)

	}
	ex, err := cmd.ExitStatus()
	if err != nil {
		t.Error("ExitStatus error: ", err)
	}
	if ex != 1 {
		t.Error("ExitStatus Expected 1, got ", ex)
	}
	_, err = cmd.Signal()
	if err != ProcessNotSignaled {
		t.Error("Signal Expected ProcessNotSignaled, got ", err)
	}
}
