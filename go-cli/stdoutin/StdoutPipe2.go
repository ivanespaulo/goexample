// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
  "bufio"
  "fmt"
  "log"
  "os/exec"
  "time"
)

func printName(jString string) {

  //cmd := exec.Command("echo", "-n", jString)
  cmdName := "docker"
  cmdArgs := []string{"build", "--no-cache=true", "--force-rm=true", "-f", "Dockerfile-Influxdb", "-t", "jeffotoni/redis:latest", "."}
  cmd := exec.Command(cmdName, cmdArgs...)

  cmdReader, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(cmdReader)
  go func() {
    for scanner.Scan() {
      //controle flux
      fmt.Println("\033[0;33m[ok]\033[0m " + scanner.Text())
    }
  }()

  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }
  if err := cmd.Wait(); err != nil {
    log.Fatal(err)
  }
}

func main() {

  printName("redis")

  // for i := 10; i < 20; i++ {
  //   go printName(`My name is Bob, I am ` + strconv.Itoa(i) + ` years old`)
  //   // Adding delay so as to see incremental output
  //   time.Sleep(60 * time.Millisecond)
  // }
  // Adding delay so as to let program complete
  // Please use channels or wait groups
  time.Sleep(100 * time.Millisecond)
}
