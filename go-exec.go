package main

import (
	"fmt"
	"log"
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

func report(c *Cmd) {
	i, err := c.Pid()
	if err != nil {
		log.Printf("Pid(%v): %v", c.Args, err)
	} else {
		log.Printf("Pid(%v): %d", c.Args, i)
	}
	s, err := c.Status()
	if err != nil {
		log.Printf("Status(%v): %v", c.Args, err)
	} else {
		log.Printf("Status(%v) => %v", c.Args, s)
	}
	b, err := c.Signaled()
	if err != nil {
		log.Printf("Signaled(%v): %v", c.Args, err)
	} else {
		log.Printf("Signaled(%v): %v", c.Args, b)
	}
	sig, err := c.Signal()
	if err != nil {
		log.Printf("Signal(%v): %v", c.Args, err)
	} else {
		log.Printf("Signal(%v): %v", c.Args, sig)
	}
	b, err = c.Exited()
	if err != nil {
		log.Printf("Exited(%v): %v", c.Args, err)
	} else {
		log.Printf("Exited(%v): %v", c.Args, b)
	}
	i, err = c.ExitStatus()
	if err != nil {
		log.Printf("ExitStatus(%v): %v", c.Args, err)
	} else {
		log.Printf("ExitStatus(%v): %v", c.Args, i)
	}
	b, err = c.Continued()
	if err != nil {
		log.Printf("Continued(%v): %v", c.Args, err)
	} else {
		log.Printf("Continued(%v): %v", c.Args, b)
	}
	b, err = c.CoreDump()
	if err != nil {
		log.Printf("CoreDump(%v): %v", c.Args, err)
	} else {
		log.Printf("CoreDump(%v): %v", c.Args, b)
	}
	b, err = c.Stopped()
	if err != nil {
		log.Printf("Stopped(%v): %v", c.Args, err)
	} else {
		log.Printf("Stopped(%v): %v", c.Args, b)
	}
	sig, err = c.StopSignal()
	if err != nil {
		log.Printf("StopSignal(%v): %v", c.Args, err)
	} else {
		log.Printf("StopSignal(%v): %v", c.Args, sig)
	}
	i, err = c.TrapCause()
	if err != nil {
		log.Printf("TrapCause(%v): %v", c.Args, err)
	} else {
		log.Printf("TrapCause(%v): %v", c.Args, i)
	}

	r, err := c.Rusage()
	if err != nil {
		log.Printf("Rusage(%v): %v", c.Args, err)
	} else {
		log.Printf("Rusage(%v) => %v", c.Args, r)
	}
	d, err := c.Utime()
	if err != nil {
		log.Printf("Utime(%v): %v", c.Args, err)
	} else {
		log.Printf("Utime(%v) => %v", c.Args, d)
	}
	d, err = c.Stime()
	if err != nil {
		log.Printf("Stime(%v): %v", c.Args, err)
	} else {
		log.Printf("Stime(%v) => %v", c.Args, d)
	}
	i6, err := c.MaxRSS()
	if err != nil {
		log.Printf("MaxRSS(%v): %v", c.Args, err)
	} else {
		log.Printf("MaxRSS(%v): %v", c.Args, i6)
	}
	i6, err = c.IxRSS()
	if err != nil {
		log.Printf("IxRSS(%v): %v", c.Args, err)
	} else {
		log.Printf("IxRSS(%v): %v", c.Args, i6)
	}
	i6, err = c.IdRSS()
	if err != nil {
		log.Printf("IdRSS(%v): %v", c.Args, err)
	} else {
		log.Printf("IdRSS(%v): %v", c.Args, i6)
	}
	i6, err = c.IsRSS()
	if err != nil {
		log.Printf("IsRSS(%v): %v", c.Args, err)
	} else {
		log.Printf("IsRSS(%v): %v", c.Args, i6)
	}
	i6, err = c.MinFlt()
	if err != nil {
		log.Printf("MinFlt(%v): %v", c.Args, err)
	} else {
		log.Printf("MinFlt(%v): %v", c.Args, i6)
	}
	i6, err = c.MajFlt()
	if err != nil {
		log.Printf("MajFlt(%v): %v", c.Args, err)
	} else {
		log.Printf("MajFlt(%v): %v", c.Args, i6)
	}
	i6, err = c.MajFlt()
	if err != nil {
		log.Printf("MajFlt(%v): %v", c.Args, err)
	} else {
		log.Printf("MajFlt(%v): %v", c.Args, i6)
	}
	i6, err = c.NSwap()
	if err != nil {
		log.Printf("NSwap(%v): %v", c.Args, err)
	} else {
		log.Printf("NSwap(%v): %v", c.Args, i6)
	}
	i6, err = c.InBlock()
	if err != nil {
		log.Printf("InBlock(%v): %v", c.Args, err)
	} else {
		log.Printf("InBlock(%v): %v", c.Args, i6)
	}
	i6, err = c.OuBlock()
	if err != nil {
		log.Printf("OuBlock(%v): %v", c.Args, err)
	} else {
		log.Printf("OuBlock(%v): %v", c.Args, i6)
	}
	i6, err = c.MsgSnd()
	if err != nil {
		log.Printf("MsgSnd(%v): %v", c.Args, err)
	} else {
		log.Printf("MsgSnd(%v): %v", c.Args, i6)
	}
	i6, err = c.MsgRcv()
	if err != nil {
		log.Printf("MsgRcv(%v): %v", c.Args, err)
	} else {
		log.Printf("MsgRcv(%v): %v", c.Args, i6)
	}
	i6, err = c.NSignals()
	if err != nil {
		log.Printf("NSignals(%v): %v", c.Args, err)
	} else {
		log.Printf("NSignals(%v): %v", c.Args, i6)
	}
	i6, err = c.NVCSw()
	if err != nil {
		log.Printf("NVCSw(%v): %v", c.Args, err)
	} else {
		log.Printf("NVCSw(%v): %v", c.Args, i6)
	}
	i6, err = c.NIvCSw()
	if err != nil {
		log.Printf("NIvCSw(%v): %v", c.Args, err)
	} else {
		log.Printf("NIvCSw(%v): %v", c.Args, i6)
	}

	fmt.Println()
}

func main() {
	for _, c := range []*Cmd{
		Command("ls", "-l", "-t"),
		Command("false"),
		Command("sleep", "1"),
	} {
		log.Println("Before Start")
		report(c)
		err := c.Start()
		if err != nil {
			log.Fatalf("Error starting %v: %v", c.Args, err)
		}

		log.Println("Before Wait")
		report(c)

		err = c.Wait()
		if err != nil {
			log.Fatalf("Error waiting for %v: %v", c.Args, err)
		}

		log.Println("After Wait")
		report(c)

		fmt.Println()
	}
}
