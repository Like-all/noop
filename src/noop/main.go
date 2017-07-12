package main

import "fmt"
import "time"

func main() {
    config = loadConfig()
    fmt.Println(config.Section("Noop").GetKey("EntryPoint"))
    for {
        fmt.Println("This is noop")
        time.Sleep(5000 * time.Millisecond)
    }
}
