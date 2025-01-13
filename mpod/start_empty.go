package main

import (
	"os"
	"os/exec"
	"syscall"
)

func StartEmpty() {
	// Run clone itself, however it will be in new namespaces
	cmd := exec.Command("bash")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func containerInit() {
	syscall.Sethostname([]byte("container"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
