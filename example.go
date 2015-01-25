package main

import (
	"cmd"
	"fmt"
	"log"
)

func report(c *cmd.Cmd) {
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
	for _, c := range []*cmd.Cmd{
		cmd.Command("ls", "-l", "-t"),
		cmd.Command("false"),
		cmd.Command("sleep", "1"),
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
