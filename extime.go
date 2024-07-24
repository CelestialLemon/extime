package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runCommand(args []string) {
	cmd := exec.Command("cmd", args...);

	output, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Println("ERROR: ", err);
		return ;
	}

	fmt.Println("OUTPUT: ", string(output[:]));
	return;
}

func main() {
	args := []string{}
	args = append(args, "/C");
	args = append(args, os.Args[1:]...);

	start := time.Now();
	runCommand(args);
	end := time.Now();

	elapsed := end.Sub(start);
	fmt.Println("Time: ", elapsed.Milliseconds(), "ms");

	return;
}