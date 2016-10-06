package main

import "fmt"
import "time"

func main() {
    for {
        fmt.Println("This is noop")
        time.Sleep(5000 * time.Millisecond)
    }
}
