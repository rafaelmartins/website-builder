package exec

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Run(cmd *exec.Cmd) (int, error) {
	if cmd == nil {
		return 0, errors.New("exec: command not defined")
	}

	// make sure that we are inheriting the default environment
	if cmd.Env != nil {
		cmd.Env = append(os.Environ(), cmd.Env...)
	}

	// attach streams
	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stdout // not a typo. pipe everything to stdout
	}

	// run command
	log.Printf("====> command: %q", cmd.Args)
	err := cmd.Run()
	if err == nil {
		log.Printf("====> success")
		return 0, nil
	}

	// get status code
	if msg, ok := err.(*exec.ExitError); ok {
		if ws, ok := msg.Sys().(syscall.WaitStatus); ok {
			status := ws.ExitStatus()
			log.Printf("====> failure: %d", status)
			return status, err
		}
	}

	log.Printf("====> failure")
	return 0, err
}
