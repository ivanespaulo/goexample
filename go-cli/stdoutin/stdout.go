// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("ls", "-lah")

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	// err = cmd.Wait()
	// //err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}

	slurp, _ := ioutil.ReadAll(stdoutIn)
	fmt.Printf("Jeff:: %s\n", slurp)
}
