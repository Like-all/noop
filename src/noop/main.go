package main

import "fmt"
import "time"
import "github.com/davecgh/go-spew/spew"

func main() {
    config, _ := loadConfig()
    spew.Dump(config)
    fmt.Println("Entry point: " + config.Noop.EntryPoint)
    for {
        fmt.Println("This is noop")
        time.Sleep(5000 * time.Millisecond)
    }
}
