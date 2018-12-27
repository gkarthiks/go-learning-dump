package main

import (
    "log"
    "os"
    "os/exec"
)

func main() {
    // cmd := exec.Command("python3 script.py")
    cmd := exec.Command("bash", "-c", "python3 script.py")
    // cmd := exec.Command("bash", "-c", "help")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    log.Println(cmd.Run())
}
