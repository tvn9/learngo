// Calling and external process
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func run() {

	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Proccess run successful with output:")
		fmt.Println(out.String())
	}
}
