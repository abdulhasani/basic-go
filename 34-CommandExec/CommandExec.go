package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output))
}
