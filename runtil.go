package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func stop(d time.Duration, p *os.Process) {
	time.Sleep(d)
	log.Printf("Timeout!\n")
	p.Signal(syscall.SIGTERM)
}

func main() {
	// check the arguments
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <end time> <command>\n", os.Args[0])
	}

	// parse the end time
	endtime, err := time.Parse("15:04", os.Args[1])
	if err != nil {
		log.Fatalf("Invalid end time (%v) Required format: 15:04\n", err)
	}
	now := time.Now()
	y, m, d := now.Date()
	l := now.Location()
	hh, mm, ss := endtime.Clock()
	et := time.Date(y, m, d, hh, mm, ss, 0, l)
	if et.Before(now) {
		et = et.Add(24 * time.Hour) // ends after midnight
	}
	runtime := et.Sub(now)
	log.Printf("Allocated time: %v\n", runtime)

	// prepare the command for execution
	c := os.Args[2:]
	cmd := exec.Command(c[0], c[1:]...)

	// redirect the outputs
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	// execute the command
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// goroutine to stop the command at the asked time
	go stop(runtime, cmd.Process)

	// wait for the command to finish
	log.Printf("Waiting for command to finish…")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
