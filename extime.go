package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runCommand(args []string) (error) {
	cmd := exec.Command("cmd", args...);

	stdout, err := cmd.StdoutPipe();
	if err != nil {
		return err;
	}

	stderr, err := cmd.StderrPipe();
	if err != nil {
		return err;
	}

	if err := cmd.Start(); err != nil {
		return err;
	}

	readAndPrint := func(pipe *bufio.Reader) {
		for {
			line, err := pipe.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(line);
		}
	}

	stdoutReader := bufio.NewReader(stdout);
	stderrReader := bufio.NewReader(stderr);

	go readAndPrint(stdoutReader);
	go readAndPrint(stderrReader);

	if err := cmd.Wait(); err != nil {
		return err;
	}

	return nil;
}

func main() {
	args := []string{}
	args = append(args, "/C");
	args = append(args, os.Args[1:]...);

	start := time.Now();
	runCommand(args);
	end := time.Now();

	elapsed := end.Sub(start);
	fmt.Println("\nextime: ", elapsed.Milliseconds(), "ms");

	return;
}