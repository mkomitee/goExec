package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

var (
	ProcessNotStarted  = fmt.Errorf("exec: not started")
	ProcessNotFinished = fmt.Errorf("exec: not finished")
	ProcessNotSignaled = fmt.Errorf("exec: not signaled")
	ProcessNotStopped  = fmt.Errorf("exec: not stopped")
	ProcessNotExited   = fmt.Errorf("exec: not exited")
	ProcessNotTrapped  = fmt.Errorf("exec: not trapped")
	StdinSet           = fmt.Errorf("exec: stdin already set")
	StdoutSet          = fmt.Errorf("exec: stdout already set")
	StderrSet          = fmt.Errorf("exec: stderr already set")
)

type Cmd struct {
	status   *syscall.WaitStatus
	rusage   *syscall.Rusage
	finished bool
	*exec.Cmd
}

func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		new(syscall.WaitStatus),
		new(syscall.Rusage),
		false,
		exec.Command(name, arg...),
	}
}

func (c *Cmd) Communicate(stdin string) (string, string, error) {
	if c.Stdin != nil {
		return "", "", StdinSet
	}
	if c.Stdout != nil {
		return "", "", StdoutSet
	}
	if c.Stderr != nil {
		return "", "", StderrSet
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdin = bytes.NewBufferString(stdin)
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	return stdout.String(), stderr.String(), err
}

func (c *Cmd) Run() error {
	if err := c.Start(); err != nil {
		return err
	}
	return c.wait4(0)
}

func (c *Cmd) Wait() error {
	return c.wait4(0)
}

func (c *Cmd) Pid() (int, error) {
	if c.Process == nil {
		return -1, ProcessNotStarted
	}
	return c.Process.Pid, nil
}

func (c *Cmd) Status() (*syscall.WaitStatus, error) {
	if c.Process == nil {
		return c.status, ProcessNotStarted
	}
	if !c.finished {
		return c.status, ProcessNotFinished
	}
	return c.status, nil
}

func (c *Cmd) Signaled() (bool, error) {
	s, err := c.Status()
	if err != nil {
		return false, err
	}
	return s.Signaled(), nil
}

func (c *Cmd) Signal() (syscall.Signal, error) {
	s, err := c.Status()
	if err != nil {
		return -1, err
	}
	if !s.Signaled() {
		return -1, ProcessNotSignaled
	}
	return s.Signal(), nil
}

func (c *Cmd) Exited() (bool, error) {
	s, err := c.Status()
	if err != nil {
		return false, err
	}
	return s.Exited(), nil
}

func (c *Cmd) ExitStatus() (int, error) {
	s, err := c.Status()
	if err != nil {
		return -1, err
	}
	if !s.Exited() {
		return -1, ProcessNotExited
	}
	return s.ExitStatus(), nil
}

func (c *Cmd) Continued() (bool, error) {
	s, err := c.Status()
	if err != nil {
		return false, err
	}
	return s.Continued(), nil
}

func (c *Cmd) CoreDump() (bool, error) {
	s, err := c.Status()
	if err != nil {
		return false, err
	}
	return s.CoreDump(), nil
}

func (c *Cmd) Stopped() (bool, error) {
	s, err := c.Status()
	if err != nil {
		return false, err
	}
	return s.Stopped(), nil
}

func (c *Cmd) StopSignal() (syscall.Signal, error) {
	s, err := c.Status()
	if err != nil {
		return -1, err
	}
	if !s.Stopped() {
		return -1, ProcessNotStopped
	}
	return s.StopSignal(), nil
}

func (c *Cmd) TrapCause() (int, error) {
	s, err := c.Status()
	if err != nil {
		return -1, err
	}
	if !s.Stopped() {
		return -1, ProcessNotStopped
	}
	if s.TrapCause() == -1 {
		return -1, ProcessNotTrapped
	}

	return s.TrapCause(), nil
}

func (c *Cmd) Rusage() (*syscall.Rusage, error) {
	if c.Process == nil {
		return c.rusage, ProcessNotStarted
	}
	if !c.finished {
		return c.rusage, ProcessNotFinished
	}
	return c.rusage, nil
}

func (c *Cmd) Utime() (time.Duration, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return time.Duration(r.Utime.Nano()), nil
}

func (c *Cmd) Stime() (time.Duration, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return time.Duration(r.Stime.Nano()), nil
}

func (c *Cmd) MaxRSS() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Maxrss, nil
}

func (c *Cmd) IxRSS() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Ixrss, nil
}

func (c *Cmd) IdRSS() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Idrss, nil
}

func (c *Cmd) IsRSS() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Isrss, nil
}

func (c *Cmd) MinFlt() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Minflt, nil
}

func (c *Cmd) MajFlt() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Majflt, nil
}

func (c *Cmd) NSwap() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Nswap, nil
}

func (c *Cmd) InBlock() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Inblock, nil
}

func (c *Cmd) OuBlock() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Oublock, nil
}

func (c *Cmd) MsgSnd() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Msgsnd, nil
}

func (c *Cmd) MsgRcv() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Msgrcv, nil
}

func (c *Cmd) NSignals() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Nsignals, nil
}

func (c *Cmd) NVCSw() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Nvcsw, nil
}

func (c *Cmd) NIvCSw() (int64, error) {
	r, err := c.Rusage()
	if err != nil {
		return 0, err
	}
	return r.Nivcsw, nil
}

func (c *Cmd) wait4(options int) error {
	pid, err := c.Pid()
	if err != nil {
		return ProcessNotStarted
	}
	_, err = syscall.Wait4(pid, c.status, options, c.rusage)
	if err != nil {
		return err
	}
	c.finished = true
	return nil
}
