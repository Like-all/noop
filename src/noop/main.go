package main

import "fmt"
import "time"

func main() {
    config, _ := loadConfig()
    fmt.Println(config.Noop.EntryPoint)
    for {
        fmt.Println("This is noop")
        time.Sleep(5000 * time.Millisecond)
    }
}
