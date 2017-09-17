package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Raspberry PI Spectrometer")
	fileName := time.Now().Format("2006-01-02_15:04::05") + ".png"
	cmd := exec.Command("raspistill", "-o", fileName)
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
}
